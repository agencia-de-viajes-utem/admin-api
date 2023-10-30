package models

// Ciudad representa la estructura de la Tabla Ciudades
// concatenada con la Tabla Paises
type Ciudad struct {
	ID         int    `json:"id"`
	Nombre     string `json:"nombre"`
	PaisID     int    `json:"pais_id"`
	NombrePais string `json:"nombre_pais"`
	AbrevPais  string `json:"abrev_pais"`
	Imagenes   string `json:"imagenes"`
}
