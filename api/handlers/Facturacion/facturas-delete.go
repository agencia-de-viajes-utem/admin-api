package handlers

import (
	"admin/api/utils"
	"encoding/json"
	"log"
	"net/http"
)

func DeleteFactura(w http.ResponseWriter, r *http.Request) {
	// Obtener el ID de factura de los parámetros de la solicitud
	idFactura := r.FormValue("id_factura")

	// Validar que se proporcionó un ID de factura
	if idFactura == "" {
		log.Printf("[%d] Se requiere el parámetro 'id_factura'", http.StatusBadRequest)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Eliminar la factura de la base de datos
	err := deleteFactura(idFactura)
	if err != nil {
		log.Printf("[%d] Error al eliminar la factura: %v", http.StatusInternalServerError, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Responder con éxito
	response := map[string]interface{}{
		"status":  "success",
		"message": "Factura eliminada con éxito",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func deleteFactura(idFactura string) error {
	// Realizar la eliminación en la base de datos
	db, err := utils.OpenDB()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("DELETE FROM facturacion WHERE id_factura = $1", idFactura)

	return err
}
