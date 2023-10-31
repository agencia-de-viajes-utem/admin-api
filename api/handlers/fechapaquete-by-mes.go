package handlers

import (
	"admin/api/models"
	"admin/api/utils"
	"encoding/json"
	"net/http"
)

// JsonRequest estructura para el JSON esperado en la solicitud
type JsonRequest struct {
	Mes           int `json:"mes"`
	TotalPersonas int `json:"total_personas"`
	CiudadOrigen  int `json:"ciudad_origen"`
	CiudadDestino int `json:"ciudad_destino"`
}

// GetFechaPaqueteByMes obtiene información de paquetes por mes
func GetFechaPaqueteByMes(w http.ResponseWriter, r *http.Request) {
	// Decodificar el cuerpo JSON de la solicitud
	var requestData JsonRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&requestData)
	if err != nil {
		http.Error(w, "Error al decodificar JSON", http.StatusBadRequest)
		return
	}

	// Utilizar los valores de la estructura JsonRequest
	mes := requestData.Mes
	totalPersonas := requestData.TotalPersonas
	ciudadOrigen := requestData.CiudadOrigen
	ciudadDestino := requestData.CiudadDestino

	// Validar que se proporcionó un mes
	if mes == 0 {
		http.Error(w, "Se requiere el parámetro 'mes'", http.StatusBadRequest)
		return
	}

	// Hacer el fetchFechaPaqueteByUser y pasarle los valores
	fechas, err := fetchFechaPaqueteByUser(mes, totalPersonas, ciudadOrigen, ciudadDestino)
	if err != nil {
		http.Error(w, "Error al obtener las fechas de paquete", http.StatusInternalServerError)
		return
	}

	// Convertir las fechas de paquete a JSON
	fechasJSON, err := json.Marshal(fechas)
	if err != nil {
		http.Error(w, "Error al convertir a JSON", http.StatusInternalServerError)
		return
	}

	// Enviar respuesta
	w.Header().Set("Content-Type", "application/json")
	w.Write(fechasJSON)
}

func fetchFechaPaqueteByUser(mes, total_personas, ciudad_origen, ciudad_destino int) ([]models.PaquetefechaInfo, error) {
	db, err := utils.OpenDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query(`
	WITH ranked_dates AS (
		SELECT
			fechapaquete.id,
			paquete.nombre,
			COALESCE(total_personas, 0) AS total_personas,
			fechapaquete.fechainit,
			fechapaquete.fechafin,
			ciudad_origen.id AS id_ciudad_origen,
			ciudad_destino.id AS id_ciudad_destino,
			ciudad_origen.nombre AS nombre_ciudad_origen,
			ciudad_destino.nombre AS nombre_ciudad_destino,
			fechapaquete.precio_oferta_vuelo as oferta_vuelo,
			paquete.precio_vuelo,
			habitacionhotel.precio_noche,
			paquete.descripcion,
			paquete.detalles,
			paquete.imagenes,
			opcionhotel.nombre AS nombre_opcion_hotel,
			habitacionhotel.descripcion AS descripcion_habitacion,
			habitacionhotel.servicios AS servicios_habitacion,
			hotel.nombre AS nombre_hotel,
			hotel.direccion AS direccion_hotel,
			hotel.valoracion AS valoracion_hotel,
			hotel.descripcion AS descripcion_hotel,
			hotel.servicios AS servicios_hotel,
			hotel.telefono AS telefono_hotel,
			hotel.correo_electronico AS correo_electronico_hotel,
			hotel.sitio_web AS sitio_web_hotel,
			ROW_NUMBER() OVER (PARTITION BY fechapaquete.id ORDER BY fechapaquete.id) AS row_num
		FROM
			paquete
			INNER JOIN unnest(paquete.id_hh) WITH ORDINALITY t(habitacion_id, ord) ON TRUE
			INNER JOIN habitacionhotel ON t.habitacion_id = habitacionhotel.id
			INNER JOIN hotel ON habitacionhotel.hotel_id = hotel.id
			INNER JOIN opcionhotel ON habitacionhotel.opcion_hotel_id = opcionhotel.id
			INNER JOIN ciudad ciudad_origen ON paquete.id_origen = ciudad_origen.id
			INNER JOIN ciudad ciudad_destino ON paquete.id_destino = ciudad_destino.id
			INNER JOIN fechapaquete ON paquete.id = fechapaquete.id_paquete
			LEFT JOIN (
				SELECT
					paquete.id AS paquete_id,
					SUM(opcionhotel.cantidad) AS total_personas
				FROM
					paquete
					INNER JOIN unnest(paquete.id_hh) WITH ORDINALITY t(habitacion_id, ord) ON TRUE
					INNER JOIN habitacionhotel ON t.habitacion_id = habitacionhotel.id
					INNER JOIN opcionhotel ON habitacionhotel.opcion_hotel_id = opcionhotel.id
				GROUP BY
					paquete.id
			) AS subquery ON paquete.id = subquery.paquete_id
	)
	SELECT *
	FROM ranked_dates
	WHERE 
    row_num = 1 	
    AND EXTRACT(MONTH FROM fechainit) = $1
    AND EXTRACT(MONTH FROM fechafin) = $1
    AND total_personas = $2
    AND id_ciudad_origen = $3
    AND id_ciudad_destino = $4;
	`, mes, total_personas, ciudad_origen, ciudad_destino)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	paquetes := make([]models.PaquetefechaInfo, 0)

	for rows.Next() {
		var paqueteInfo models.PaquetefechaInfo
		var infoPaquete models.PaquetefechaInfoAdicional
		var hotelInfo models.HotelfechaInfo

		err := rows.Scan(
			&paqueteInfo.ID,
			&paqueteInfo.Nombre,
			&paqueteInfo.TotalPersonas,
			&paqueteInfo.FechaInit,
			&paqueteInfo.FechaFin,
			&paqueteInfo.IdOrigen,
			&paqueteInfo.IdDestino,
			&paqueteInfo.NombreCiudadOrigen,
			&paqueteInfo.NombreCiudadDestino,
			&paqueteInfo.PrecioOfertaVuelo,
			&paqueteInfo.PrecioVuelo,
			&paqueteInfo.PrecioNoche,
			&paqueteInfo.Descripcion,
			&paqueteInfo.Detalles,
			&paqueteInfo.Imagenes,
			&infoPaquete.NombreOpcionHotel,
			&infoPaquete.DescripcionHabitacion,
			&infoPaquete.ServiciosHabitacion,
			&hotelInfo.NombreHotel,
			&hotelInfo.DireccionHotel,
			&hotelInfo.ValoracionHotel,
			&hotelInfo.DescripcionHotel,
			&hotelInfo.ServiciosHotel,
			&hotelInfo.TelefonoHotel,
			&hotelInfo.CorreoElectronico,
			&hotelInfo.SitioWeb,
			&infoPaquete.RowNum,
		)

		if err != nil {
			return nil, err
		}

		infoPaquete.HotelInfo = hotelInfo
		paqueteInfo.InfoPaquetefecha = infoPaquete

		paquetes = append(paquetes, paqueteInfo)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return paquetes, nil
}
