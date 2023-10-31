package handlers

import (
	"admin/api/utils"
	"encoding/json"
	"log"
	"net/http"
)

// CiudadDeleteRequest representa la estructura para eliminar una ciudad por su ID
type CiudadDeleteRequest struct {
	ID int `json:"id"`
}

func DeleteCiudad(w http.ResponseWriter, r *http.Request) {
	var deleteRequest CiudadDeleteRequest

	if err := json.NewDecoder(r.Body).Decode(&deleteRequest); err != nil {
		log.Printf("[%d] Error al decodificar los datos de eliminación: %v", http.StatusBadRequest, err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := deleteCiudad(deleteRequest.ID); err != nil {
		log.Printf("[%d] Error al eliminar la ciudad: %v", http.StatusInternalServerError, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"status":  "success",
		"message": "Ciudad eliminada con éxito",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func deleteCiudad(id int) error {
	// Realizar la eliminación en la base de datos
	db, err := utils.OpenDB()
	if err != nil {
		return err
	}
	defer db.Close()

	// Realizar la consulta DELETE en la base de datos utilizando el ID
	_, err = db.Exec("DELETE FROM ciudad WHERE id = $1", id)

	return err
}
