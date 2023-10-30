package handlers

import (
	"admin/api/utils"
	"net/http"
)

func DeleteFactura(w http.ResponseWriter, r *http.Request) {
	// Obtener el ID de factura de los parámetros de la solicitud
	idFactura := r.FormValue("id_factura")

	// Validar que se proporcionó un ID de factura
	if idFactura == "" {
		handleError(w, "Se requiere el parámetro 'id_factura'", http.StatusBadRequest, nil)
		return
	}

	// Eliminar la factura de la base de datos
	err := deleteFactura(idFactura)
	if err != nil {
		handleError(w, "Error al eliminar la factura", http.StatusInternalServerError, err)
		return
	}

	// Responder con éxito
	w.WriteHeader(http.StatusOK)
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
