package handlers

import (
	"admin/api/utils"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

// FacturaUpdate representa la estructura para la actualización de la factura
type FacturaUpdate struct {
	IDFactura     int     `json:"idFactura"`
	MontoTotal    float64 `json:"monto_total"`
	FechaCreacion string  `json:"fecha_creacion"`
	FkReserva     int     `json:"fk_reserva"`
}

func UpdateFactura(w http.ResponseWriter, r *http.Request) {
	// Obtener los datos actualizados de la factura desde el cuerpo de la solicitud
	var facturaUpdate FacturaUpdate

	err := json.NewDecoder(r.Body).Decode(&facturaUpdate)
	if err != nil {
		log.Printf("[%d] Error al decodificar los datos de actualización: %v", http.StatusBadRequest, err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Actualizar la factura en la base de datos
	err = updateFactura(facturaUpdate.IDFactura, facturaUpdate.MontoTotal, facturaUpdate.FechaCreacion, facturaUpdate.FkReserva)
	if err != nil {
		log.Printf("[%d] Error al actualizar la factura: %v", http.StatusInternalServerError, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Responder con éxito
	w.WriteHeader(http.StatusOK)
}

func updateFactura(idFactura int, montoTotal float64, fechaCreacion string, fkReserva int) error {
	// Convertir la fecha de creación a formato de tiempo
	fechaCreacionTime, err := time.Parse("2006-01-02 15:04:05", fechaCreacion)
	if err != nil {
		return err
	}

	// Realizar la actualización en la base de datos
	db, err := utils.OpenDB()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec(`
		UPDATE facturacion
		SET monto_total = $1, fecha_creacion = $2, fk_reserva = $3
		WHERE id_factura = $4
	`, montoTotal, fechaCreacionTime, fkReserva, idFactura)

	return err
}
