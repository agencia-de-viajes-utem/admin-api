package handlers

import (
	"admin/api/models"
	"encoding/json"
	"net/http"
)

func GetFacturasByUsuarios(w http.ResponseWriter, r *http.Request) {
	// Obtener el ID de usuario del formulario o de los parámetros de la solicitud
	idUsuario := r.FormValue("id_usuario")

	// Validar que se proporcionó un ID de usuario
	if idUsuario == "" {
		handleError(w, "Se requiere el parámetro 'id_usuario'", http.StatusBadRequest, nil)
		return
	}

	// Obtener todas las reservas
	reservas, err := fetchReservas()
	if err != nil {
		handleError(w, "Error al obtener las reservas", http.StatusInternalServerError, err)
		return
	}

	// Filtrar las reservas por el ID de usuario
	reservasFiltradas := filterFacturasByUsuarios(reservas, idUsuario)

	// Convertir a JSON y responder
	reservasJSON, err := json.Marshal(reservasFiltradas)
	if err != nil {
		handleError(w, "Error al convertir a JSON", http.StatusInternalServerError, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(reservasJSON)
}

func filterFacturasByUsuarios(reservas []models.Factura, idUsuario string) []models.Factura {
	var reservasFiltradas []models.Factura

	// Filtrar las reservas por el ID de usuario
	for _, reserva := range reservas {
		if reserva.IDUsuario == idUsuario {
			reservasFiltradas = append(reservasFiltradas, reserva)
		}
	}

	return reservasFiltradas
}
