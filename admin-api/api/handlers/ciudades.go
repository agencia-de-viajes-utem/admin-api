package handlers

import (
	"admin/api/models"
	"admin/api/utils"
	"encoding/json"
	"net/http"
)

// GetAllCiudades obtiene todas las ciudades y las devuelve como JSON.
func GetAllCiudades(w http.ResponseWriter, r *http.Request) {
	ciudades, err := fetchCiudades()
	if err != nil {
		handleError(w, "Error al obtener las ciudades", http.StatusInternalServerError, err)
		return
	}

	ciudadesJSON, err := json.Marshal(ciudades)
	if err != nil {
		handleError(w, "Error al convertir a JSON", http.StatusInternalServerError, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(ciudadesJSON)
}

func fetchCiudades() ([]models.Ciudad, error) {
	db, err := utils.OpenDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query(`
		SELECT ciudad.*, pais.nombre as nombre_pais, pais.abreviacion as abrev_pais
		FROM ciudad
		INNER JOIN pais ON ciudad.pais_id = pais.id;
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	ciudades := make([]models.Ciudad, 0)

	for rows.Next() {
		ciudad := models.Ciudad{}
		err := rows.Scan(
			&ciudad.ID,
			&ciudad.Nombre,
			&ciudad.PaisID,
			&ciudad.NombrePais,
			&ciudad.AbrevPais,
			&ciudad.Imagenes,
		)
		if err != nil {
			return nil, err
		}
		ciudades = append(ciudades, ciudad)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return ciudades, nil
}
