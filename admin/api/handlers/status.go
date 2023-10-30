package handlers

import (
	"admin/api/utils"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// ErrorResponse representa la estructura de una respuesta de error.
type ErrorResponse struct {
	Message string `json:"message"`
	Error   string `json:"error,omitempty"`
}

// handleError maneja los errores y envía una respuesta HTTP con el mensaje de error correspondiente.
func handleError(w http.ResponseWriter, errMsg string, statusCode int, err error) {
	log.Printf("[%d] %s: %v", statusCode, errMsg, err)

	response := ErrorResponse{
		Message: errMsg,
		Error:   err.Error(),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Println("Error al escribir la respuesta de error:", err)
	}
}

// handleError maneja los errores y envía una respuesta HTTP con el mensaje de error correspondiente.
func handleError2(w http.ResponseWriter, err error, status int) {
	fmt.Println("Error:", err)
	w.WriteHeader(status)
}

// HomeHandler maneja las solicitudes a la ruta "/"
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// Abrir la conexión a la base de datos
	db, err := utils.OpenDB()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}
	defer db.Close()

	// Intentar hacer ping a la base de datos
	err = db.Ping()
	if err != nil {
		// Si hay un error al hacer ping, devolver código 403 (Forbidden)
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("Forbidden"))
		return
	}

	// Si se hizo ping con éxito, devolver código 200 (OK)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
