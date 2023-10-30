package handlers

import (
	"admin/api/models"
	"admin/api/utils"
	"encoding/json"
	"net/http"
)

// GetAllPaises obtiene todos los países y los devuelve como JSON.
func GetAllPaises(w http.ResponseWriter, r *http.Request) {
	paises, err := fetchPaises()
	if err != nil {
		handleError(w, "Error al obtener los países", http.StatusInternalServerError, err)
		return
	}

	paisesJSON, err := json.Marshal(paises)
	if err != nil {
		handleError(w, "Error al convertir a JSON", http.StatusInternalServerError, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(paisesJSON)
}

func fetchPaises() ([]models.Pais, error) {
	db, err := utils.OpenDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query(`
		SELECT *
		FROM pais;
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	paises := make([]models.Pais, 0)

	for rows.Next() {
		pais := models.Pais{}
		err := rows.Scan(
			&pais.ID,
			&pais.Nombre,
			&pais.Abreviacion,
			&pais.Imagenes,
		)
		if err != nil {
			return nil, err
		}
		paises = append(paises, pais)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return paises, nil
}
