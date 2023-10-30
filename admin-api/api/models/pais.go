package models

// Pais representa la estructura de la Tabla Paises

type Pais struct {
	ID          int    `json:"id"`
	Nombre      string `json:"nombre"`
	Abreviacion string `json:"abreviacion"`
	Imagenes    string `json:"imagenes"`
}
