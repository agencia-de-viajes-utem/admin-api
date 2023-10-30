package handlers

import (
	"admin/api/models"
	"admin/api/utils"
	"encoding/json"
	"net/http"
)

// GetAllAeropuertos obtiene todos los aeropuertos y los devuelve como JSON.
func GetAllAeropuertos(w http.ResponseWriter, r *http.Request) {
	aeropuertos, err := fetchAeropuertos()
	if err != nil {
		handleError(w, "Error al obtener los aeropuertos", http.StatusInternalServerError, err)
		return
	}

	aeropuertosJSON, err := json.Marshal(aeropuertos)
	if err != nil {
		handleError(w, "Error al convertir a JSON", http.StatusInternalServerError, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(aeropuertosJSON)
}

func fetchAeropuertos() ([]models.Aeropuerto, error) {
	db, err := utils.OpenDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query(`
		SELECT aeropuerto.*, ciudad.nombre as nombre_ciudad,ciudad.pais_id as id_pais, pais.nombre as nombre_pais, pais.abreviacion as abrev_pais
		FROM aeropuerto
		INNER JOIN ciudad ON aeropuerto.ciudad_id = ciudad.id
		INNER JOIN pais ON ciudad.pais_id = pais.id;
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	aeropuertos := make([]models.Aeropuerto, 0)

	for rows.Next() {
		aeropuerto := models.Aeropuerto{}
		err := rows.Scan(
			&aeropuerto.ID,
			&aeropuerto.Nombre,
			&aeropuerto.CiudadID,
			&aeropuerto.NombreCiudad,
			&aeropuerto.PaisID,
			&aeropuerto.NombrePais,
			&aeropuerto.AbrevPais,
		)
		if err != nil {
			return nil, err
		}
		aeropuertos = append(aeropuertos, aeropuerto)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return aeropuertos, nil
}
