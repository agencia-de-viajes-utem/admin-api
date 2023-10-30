package models

// Aeropuerto representa la estructura de la Tabla Aeropuertos
// concatenada con la Tabla Ciudades y la Tabla Paises
type Aeropuerto struct {
	ID           int    `json:"id"`
	Nombre       string `json:"nombre"`
	CiudadID     int    `json:"ciudad_id"`
	NombreCiudad string `json:"nombre_ciudad"`
	PaisID       int    `json:"pais_id"`
	NombrePais   string `json:"nombre_pais"`
	AbrevPais    string `json:"abrev_pais"`
}
