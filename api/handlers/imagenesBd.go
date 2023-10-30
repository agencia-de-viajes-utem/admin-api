package handlers

import (
	"admin/api/models"
	"admin/api/utils"
	"encoding/json"
	"net/http"
)

func ImagenesBd(w http.ResponseWriter, r *http.Request) {
	imagenes, err := FetchBdInfoImagenes()
	if err != nil {
		handleError(w, "Error al obtener las imágenes", http.StatusInternalServerError, err)
		return
	}

	imagenesJSON, err := json.Marshal(imagenes)
	if err != nil {
		handleError(w, "Error al convertir a JSON", http.StatusInternalServerError, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(imagenesJSON)
}

// FetchBdInfoImagenes obtiene la información de las imágenes desde la base de datos.
func FetchBdInfoImagenes() ([]models.Imagen, error) {
	imagenes, err := fetchImagenesFromDb()
	if err != nil {
		return nil, err
	}

	return imagenes, nil
}

func fetchImagenesFromDb() ([]models.Imagen, error) {
	db, err := utils.OpenDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query(`
		SELECT id, alt, descripcion, fecha_creacion
		FROM imagen;
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	imagenes := make([]models.Imagen, 0)

	for rows.Next() {
		imagen := models.Imagen{}
		err := rows.Scan(&imagen.ID, &imagen.Alt, &imagen.Descripcion, &imagen.FechaCreacion)
		if err != nil {
			return nil, err
		}
		imagenes = append(imagenes, imagen)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return imagenes, nil
}
