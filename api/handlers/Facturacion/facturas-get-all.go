package handlers

import (
	"admin/api/models"
	"admin/api/utils"
	"encoding/json"
	"log"
	"net/http"
)

func GetAllFacturas(w http.ResponseWriter, r *http.Request) {
	facturas, err := fetchFacturas()
	if err != nil {
		log.Printf("[%d] Error al obtener las facturas: %v", http.StatusInternalServerError, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	facturasJSON, err := json.Marshal(facturas)
	if err != nil {
		log.Printf("[%d] Error al convertir a JSON: %v", http.StatusInternalServerError, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(facturasJSON)
}

func fetchFacturas() ([]models.Factura, error) {
	db, err := utils.OpenDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query(`SELECT 
    id_factura,
	monto_total,
    reserva.id AS id_reserva,
	reserva.estado as estado_reserva,
	reserva.id_usuario,
	reserva.pasajeros,
    JSON_BUILD_OBJECT(
        'id_paquete', id_paquete,
        'id_fp', fp.id, -- Corregido para incluir la columna id_fp
        'precio_oferta_vuelo', fp.precio_oferta_vuelo,
        'precio_vuelo', paquete.precio_vuelo,
        'precio_total_noches', SUM(habitacionhotel.precio_noche),
        'capacidad_total', SUM(opcionhotel.cantidad),
        'fecha_inicial', fp.fechainit,
        'fecha_final', fp.fechafin,
        'ciudad_origen', ciudad_origen.nombre,
        'ciudad_destino', ciudad_destino.nombre,
        'pais_origen', pais_origen.nombre,
        'pais_origen_abrev', pais_origen.abreviacion,
        'pais_destino', pais_destino.nombre,
        'pais_destino_abrev', pais_destino.abreviacion,
        'paquete_descripcion', paquete.descripcion,
        'paquete_detalle', paquete.detalles,
        'paquete_imagenes', paquete.imagenes,
		    'id_hh', paquete.id_hh,
        'info_hotel', JSON_BUILD_OBJECT(
            'hotel_id', hotel.id,
            'hotel_nombre', hotel.nombre,
            'hotel_ciudad', ciudad_hotel.nombre,
            'hotel_direccion', hotel.direccion,
            'hotel_valoracion', hotel.valoracion,
            'hotel_descripcion', hotel.descripcion,
            'hotel_servicios', hotel.servicios,
            'hotel_telefono', hotel.telefono,
            'hotel_email', hotel.correo_electronico,
            'hotel_sitio_web', hotel.sitio_web,
            'hotel_imagenes', hotel.imagenes,
            'info_habitaciones', JSON_AGG(
                JSON_BUILD_OBJECT(
					'habitacion_id', habitacionhotel.id,
                    'hh_tipo', opcionhotel.nombre,
                    'hh_descripcion', habitacionhotel.descripcion,
                    'hh_servicios', habitacionhotel.servicios,
                    'hh_imagenes', habitacionhotel.imagenes
                )
            )
        )
    ) AS info_paquete
FROM 
    facturacion
INNER JOIN reserva ON reserva.id = fk_reserva
INNER JOIN fechapaquete AS fp ON fp.id = id_fechapaquete
INNER JOIN paquete ON paquete.id = id_paquete
INNER JOIN ciudad AS ciudad_origen ON id_origen = ciudad_origen.id
INNER JOIN ciudad AS ciudad_destino ON id_destino = ciudad_destino.id
INNER JOIN pais AS pais_origen ON ciudad_origen.pais_id = pais_origen.id
INNER JOIN pais AS pais_destino ON ciudad_destino.pais_id = pais_destino.id
INNER JOIN unnest(paquete.id_hh) WITH ORDINALITY t(habitacion_id, ord) ON TRUE
INNER JOIN habitacionhotel ON t.habitacion_id = habitacionhotel.id
INNER JOIN opcionhotel ON habitacionhotel.opcion_hotel_id = opcionhotel.id
INNER JOIN hotel ON habitacionhotel.hotel_id = hotel.id
INNER JOIN ciudad AS ciudad_hotel ON hotel.ciudad_id = ciudad_hotel.id
GROUP BY
	hotel.id,
    paquete.id_hh,
    id_factura,
    id_paquete,
    fp.id, -- Agregado para incluir la columna id_fp
    reserva.id,
    fp.precio_oferta_vuelo,
    paquete.precio_vuelo,
    fp.fechainit,
    fp.fechafin,
    ciudad_origen.nombre,
    ciudad_destino.nombre,
    pais_origen.nombre,
    pais_origen.abreviacion,
    pais_destino.nombre,
    pais_destino.abreviacion,
    paquete.descripcion,
    paquete.detalles,
    paquete.imagenes,
    habitacionhotel.hotel_id,
    hotel.nombre,
    ciudad_hotel.nombre,
    hotel.direccion,
    hotel.valoracion,
    hotel.descripcion,
    hotel.servicios,
    hotel.telefono,
    hotel.correo_electronico,
    hotel.sitio_web,
    hotel.imagenes;
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	facturas := make([]models.Factura, 0)

	for rows.Next() {
		var id_factura int
		var montoTotal float64
		var idReserva int
		var estadoReserva string
		var idUsuario string
		var pasajerosJSON string
		var infoPaqueteStr string

		err := rows.Scan(&id_factura, &montoTotal, &idReserva, &estadoReserva, &idUsuario, &pasajerosJSON, &infoPaqueteStr)
		if err != nil {
			return nil, err
		}

		infoPaquete, err := models.ParseInfoPaquete(infoPaqueteStr)
		if err != nil {
			return nil, err
		}

		// Parsear los pasajeros
		var pasajerosObj models.Pasajeros
		err = json.Unmarshal([]byte(pasajerosJSON), &pasajerosObj)
		if err != nil {
			return nil, err
		}

		// Acceder al slice de pasajeros
		factura := models.Factura{
			IDReserva:   idReserva,
			Estado:      estadoReserva,
			IDFactura:   id_factura,
			IDUsuario:   idUsuario,
			Pasajeros:   pasajerosObj.Pasajeros,
			InfoPaquete: *infoPaquete,
			MontoTotal:  montoTotal,
		}

		facturas = append(facturas, factura)
	}

	return facturas, nil

}
