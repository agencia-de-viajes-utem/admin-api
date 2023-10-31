package handlers

import (
	"admin/api/utils"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

// FacturaCreate representa la estructura para la creación de una nueva factura
type FacturaCreate struct {
	MontoTotal    float64 `json:"monto_total"`
	FechaCreacion string  `json:"fecha_creacion"`
	FkReserva     int     `json:"fk_reserva"`
}

func CreateFactura(w http.ResponseWriter, r *http.Request) {
	// Obtener los datos para la creación de la factura desde el cuerpo de la solicitud
	var facturaCreate FacturaCreate

	err := json.NewDecoder(r.Body).Decode(&facturaCreate)
	if err != nil {
		log.Printf("[%d] Error al decodificar los datos de creación: %v", http.StatusBadRequest, err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Crear la nueva factura en la base de datos
	err = createFactura(facturaCreate.MontoTotal, facturaCreate.FechaCreacion, facturaCreate.FkReserva)
	if err != nil {
		log.Printf("[%d] Error al crear la factura: %v", http.StatusInternalServerError, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Responder con éxito
	w.WriteHeader(http.StatusCreated)
}

func createFactura(montoTotal float64, fechaCreacion string, fkReserva int) error {
	// Convertir la fecha de creación a formato de tiempo
	fechaCreacionTime, err := time.Parse("2006-01-02 15:04:05", fechaCreacion)
	if err != nil {
		return err
	}

	// Realizar la verificación en la base de datos
	db, err := utils.OpenDB()
	if err != nil {
		return err
	}
	defer db.Close()

	// Verificar si ya existe una factura con la misma reserva
	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM facturacion WHERE fk_reserva = $1", fkReserva).Scan(&count)
	if err != nil {
		return err
	}

	// Si count es mayor que cero, significa que ya hay una factura con la misma reserva
	if count > 0 {
		return fmt.Errorf("ya existe una factura con la reserva especificada")
	}

	// Realizar la inserción en la base de datos
	_, err = db.Exec(`
		INSERT INTO facturacion (monto_total, fecha_creacion, fk_reserva)
		VALUES ($1, $2, $3)
	`, montoTotal, fechaCreacionTime, fkReserva)

	return err
}
