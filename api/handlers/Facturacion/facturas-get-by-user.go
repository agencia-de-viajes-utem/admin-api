package handlers

import (
	"admin/api/models"
	"encoding/json"
	"log"
	"net/http"
)

func GetFacturasByUsuarios(w http.ResponseWriter, r *http.Request) {
	// Obtener el ID de usuario del formulario o de los parámetros de la solicitud
	idUsuario := r.FormValue("id_usuario")

	// Validar que se proporcionó un ID de usuario
	if idUsuario == "" {
		log.Printf("[%d] Se requiere el parámetro 'id_usuario'", http.StatusBadRequest)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Obtener todas las facturas
	facturas, err := fetchFacturas()
	if err != nil {
		log.Printf("[%d] Error al obtener las facturas: %v", http.StatusInternalServerError, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Filtrar las facturas por el ID de usuario
	facturasFiltradas := filterFacturasByUsuarios(facturas, idUsuario)

	// Convertir a JSON y responder
	facturasJSON, err := json.Marshal(facturasFiltradas)
	if err != nil {
		log.Printf("[%d] Error al convertir a JSON: %v", http.StatusInternalServerError, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(facturasJSON)
}

func filterFacturasByUsuarios(facturas []models.Factura, idUsuario string) []models.Factura {
	var facturasFiltradas []models.Factura

	// Filtrar las facturas por el ID de usuario
	for _, factura := range facturas {
		if factura.IDUsuario == idUsuario {
			facturasFiltradas = append(facturasFiltradas, factura)
		}
	}

	return facturasFiltradas
}
