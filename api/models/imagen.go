package models

import (
	"encoding/json"
	"time"
)

// Imagen representa el modelo de datos para una imagen.
type Imagen struct {
	ID            string    `json:"id"`
	Alt           string    `json:"alt"`
	Descripcion   string    `json:"descripcion"`
	FechaCreacion time.Time `json:"fecha_creacion"`
}

// InfoBdImagenes representa la información de las imágenes desde la base de datos.
type InfoBdImagenes struct {
	ID            string `json:"id"`
	Alt           string `json:"alt"`
	Descripcion   string `json:"descripcion"`
	FechaCreacion string `json:"fecha_creacion"`
	UrlPublica    string `json:"url_publica"`
}

// MarshalJSON implementa la interfaz Marshaler para personalizar la serialización JSON.
func (i Imagen) MarshalJSON() ([]byte, error) {
	type Alias Imagen
	return json.Marshal(&struct {
		Alias
		FechaCreacion string `json:"fecha_creacion"`
	}{
		Alias:         (Alias)(i),
		FechaCreacion: i.FechaCreacion.Format("2006-01-02 15:04:05"),
	})
}

// UnmarshalJSON implementa la interfaz Unmarshaler para personalizar la deserialización JSON.
func (i *Imagen) UnmarshalJSON(data []byte) error {
	type Alias Imagen
	aux := &struct {
		*Alias
		FechaCreacion string `json:"fecha_creacion"`
	}{
		Alias: (*Alias)(i),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	t, err := time.Parse("2006-01-02 15:04:05", aux.FechaCreacion)
	if err != nil {
		return err
	}
	i.FechaCreacion = t
	return nil
}
