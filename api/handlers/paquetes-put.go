package handlers

import (
	"admin/api/utils"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/lib/pq"
)

// PaqueteUpdateRequest representa la estructura para actualizar un paquete
type PaqueteUpdateRequest struct {
	ID          int      `json:"id"`
	Nombre      string   `json:"nombre"`
	IdOrigen    int      `json:"id_origen"`
	IdDestino   int      `json:"id_destino"`
	Descripcion string   `json:"descripcion"`
	Detalles    string   `json:"detalles"`
	PrecioVuelo float64  `json:"precio_vuelo"`
	IdHh        []int    `json:"id_hh"`
	Imagenes    []string `json:"imagenes"`
}

func UpdatePaquete(w http.ResponseWriter, r *http.Request) {
	// Leer y decodificar el JSON content de la solicitud PUT
	var request PaqueteUpdateRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		handleError(w, "Error al decodificar los datos de actualización", http.StatusBadRequest, err)
		return
	}

	// Llamar a la función que maneja la actualización del paquete
	err = updatePaquete(request)
	if err != nil {
		handleError(w, "Error al actualizar el paquete", http.StatusInternalServerError, err)
		return
	}

	// Responder con una respuesta JSON apropiada
	response := map[string]interface{}{
		"status":  "success",
		"message": "Paquete actualizado con éxito",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func updatePaquete(request PaqueteUpdateRequest) error {
	// Realizar la actualización en la base de datos
	db, err := utils.OpenDB()
	if err != nil {
		return err
	}
	defer db.Close()

	// Preparar la consulta de actualización
	updateQuery := "UPDATE paquete SET "
	updateFields := []string{}
	queryParams := []interface{}{request.ID}

	// Verificar y agregar campos actualizables
	if request.Nombre != "" {
		updateFields = append(updateFields, "nombre = $"+strconv.Itoa(len(queryParams)+1))
		queryParams = append(queryParams, request.Nombre)
	}

	if request.IdOrigen != 0 {
		updateFields = append(updateFields, "id_origen = $"+strconv.Itoa(len(queryParams)+1))
		queryParams = append(queryParams, request.IdOrigen)
	}

	if request.IdDestino != 0 {
		updateFields = append(updateFields, "id_destino = $"+strconv.Itoa(len(queryParams)+1))
		queryParams = append(queryParams, request.IdDestino)
	}

	if request.Descripcion != "" {
		updateFields = append(updateFields, "descripcion = $"+strconv.Itoa(len(queryParams)+1))
		queryParams = append(queryParams, request.Descripcion)
	}

	if request.Detalles != "" {
		updateFields = append(updateFields, "detalles = $"+strconv.Itoa(len(queryParams)+1))
		queryParams = append(queryParams, request.Detalles)
	}

	if request.PrecioVuelo != 0 {
		updateFields = append(updateFields, "precio_vuelo = $"+strconv.Itoa(len(queryParams)+1))
		queryParams = append(queryParams, request.PrecioVuelo)
	}

	// Actualizar campos de tipo array (IdHh e Imagenes)
	if len(request.IdHh) > 0 {
		updateFields = append(updateFields, "id_hh = $"+strconv.Itoa(len(queryParams)+1))
		queryParams = append(queryParams, pq.Array(request.IdHh))
	}

	if len(request.Imagenes) > 0 {
		updateFields = append(updateFields, "imagenes = $"+strconv.Itoa(len(queryParams)+1))
		queryParams = append(queryParams, pq.Array(request.Imagenes))
	}

	// Combinar los campos actualizables en la consulta
	updateQuery += strings.Join(updateFields, ", ") + " WHERE id = $1"

	// Realizar la actualización en la base de datos
	_, err = db.Exec(updateQuery, queryParams...)

	return err
}
