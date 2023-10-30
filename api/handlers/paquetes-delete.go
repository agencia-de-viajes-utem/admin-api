package handlers

import (
	"admin/api/utils"
	"encoding/json"
	"net/http"
)

// PaqueteDeleteRequest representa la estructura para eliminar un paquete por su ID
type PaqueteDeleteRequest struct {
	ID int `json:"id"`
}

func DeletePaquete(w http.ResponseWriter, r *http.Request) {
	// Obtener los datos para eliminar el paquete desde el cuerpo de la solicitud
	var deleteRequest PaqueteDeleteRequest

	err := json.NewDecoder(r.Body).Decode(&deleteRequest)
	if err != nil {
		handleError(w, "Error al decodificar los datos de eliminación", http.StatusBadRequest, err)
		return
	}

	// Realizar la eliminación en la base de datos
	err = deletePaquete(deleteRequest.ID)
	if err != nil {
		handleError(w, "Error al eliminar el paquete", http.StatusInternalServerError, err)
		return
	}

	// Responder con éxito y proporcionar una respuesta JSON
	response := map[string]interface{}{
		"status":  "success",
		"message": "Paquete eliminado con éxito",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func deletePaquete(id int) error {
	// Realizar la eliminación en la base de datos
	db, err := utils.OpenDB()
	if err != nil {
		return err
	}
	defer db.Close()

	// Realizar la consulta DELETE en la base de datos utilizando el ID
	_, err = db.Exec("DELETE FROM paquete WHERE id = $1", id)

	return err
}
