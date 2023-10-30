package models

import "encoding/json"

// HabitacionHotel representa la información de una habitación de hotel.
type HabitacionHotel struct {
	ID          int    `json:"habitacion_id"`
	Tipo        string `json:"hh_tipo"`
	Descripcion string `json:"hh_descripcion"`
	Servicios   string `json:"hh_servicios"`
	Imagenes    string `json:"hh_imagenes"`
}

// InfoHotel representa la información de un hotel.
type InfoHotel struct {
	HotelID          int               `json:"hotel_id"`
	HotelNombre      string            `json:"hotel_nombre"`
	HotelCiudad      string            `json:"hotel_ciudad"`
	HotelDireccion   string            `json:"hotel_direccion"`
	HotelValoracion  float64           `json:"hotel_valoracion"`
	HotelDescripcion string            `json:"hotel_descripcion"`
	HotelServicios   string            `json:"hotel_servicios"`
	HotelTelefono    string            `json:"hotel_telefono"`
	HotelEmail       string            `json:"hotel_email"`
	HotelSitioWeb    string            `json:"hotel_sitio_web"`
	HotelImagenes    string            `json:"hotel_imagenes"`
	InfoHabitaciones []HabitacionHotel `json:"info_habitaciones"`
}

// Pasajero representa la información de un pasajero.
type Pasajero struct {
	Nombre    string `json:"nombre"`
	Apellidos string `json:"apellidos"`
	RUT       string `json:"rut"`
	Correo    string `json:"correo"`
	Numero    string `json:"numero"`
}

// Pasajeros representa la información de los pasajeros.
type Pasajeros struct {
	Pasajeros []Pasajero `json:"pasajeros"`
}

// Reserva representa la información de una reserva.
type Factura struct {
	IDFactura   int         `json:"id_factura"`
	MontoTotal  float64     `json:"monto_total"`
	IDReserva   int         `json:"id_reserva"`
	Estado      string      `json:"estado_reserva"`
	IDUsuario   string      `json:"id_usuario"`
	Pasajeros   []Pasajero  `json:"pasajeros"`
	InfoPaquete InfoPaquete `json:"info_paquete"`
}

// InfoPaquete representa la información de un paquete.
type InfoPaquete struct {
	PrecioOfertaVuelo  float64   `json:"precio_oferta_vuelo"`
	PrecioVuelo        float64   `json:"precio_vuelo"`
	PrecioTotalNoches  float64   `json:"precio_total_noches"`
	CapacidadTotal     int       `json:"capacidad_total"`
	FechaInicial       string    `json:"fecha_inicial"`
	FechaFinal         string    `json:"fecha_final"`
	CiudadOrigen       string    `json:"ciudad_origen"`
	CiudadDestino      string    `json:"ciudad_destino"`
	PaisOrigen         string    `json:"pais_origen"`
	PaisOrigenAbrev    string    `json:"pais_origen_abrev"`
	PaisDestino        string    `json:"pais_destino"`
	PaisDestinoAbrev   string    `json:"pais_destino_abrev"`
	PaqueteDescripcion string    `json:"paquete_descripcion"`
	PaqueteDetalle     string    `json:"paquete_detalle"`
	PaqueteImagenes    string    `json:"paquete_imagenes"`
	IDHH               []int     `json:"id_hh"`
	InfoHotel          InfoHotel `json:"info_hotel"`
}

// ParseInfoPaquete convierte una cadena JSON en una estructura InfoPaquete.
func ParseInfoPaquete(jsonStr string) (*InfoPaquete, error) {
	var infoPaquete InfoPaquete
	err := json.Unmarshal([]byte(jsonStr), &infoPaquete)
	if err != nil {
		return nil, err
	}
	return &infoPaquete, nil
}
