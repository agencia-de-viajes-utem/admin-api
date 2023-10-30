package handlers

import (
	"admin/api/utils"
	"encoding/json"
	"net/http"
)

// CiudadDeleteRequest representa la estructura para eliminar una ciudad por su ID
type CiudadDeleteRequest struct {
	ID int `json:"id"`
}

func DeleteCiudad(w http.ResponseWriter, r *http.Request) {
	// Obtener los datos para eliminar la ciudad desde el cuerpo de la solicitud
	var deleteRequest CiudadDeleteRequest

	err := json.NewDecoder(r.Body).Decode(&deleteRequest)
	if err != nil {
		handleError(w, "Error al decodificar los datos de eliminación", http.StatusBadRequest, err)
		return
	}

	// Realizar la eliminación en la base de datos
	err = deleteCiudad(deleteRequest.ID)
	if err != nil {
		handleError(w, "Error al eliminar la ciudad", http.StatusInternalServerError, err)
		return
	}

	// Responder con éxito y proporcionar una respuesta JSON
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
