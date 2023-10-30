package handlers

import (
	"admin/api/utils"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/lib/pq"
)

// CiudadUpdate representa la estructura para la actualización parcial de una ciudad
type CiudadUpdate struct {
	ID       int      `json:"id"`
	Nombre   *string  `json:"nombre,omitempty"`
	PaisID   *int     `json:"pais_id,omitempty"`
	Imagenes []string `json:"imagenes,omitempty"`
}

// UpdateCiudad actualiza una ciudad en la base de datos según los campos proporcionados
func UpdateCiudad(w http.ResponseWriter, r *http.Request) {
	// Obtener los datos actualizados de la ciudad desde el cuerpo de la solicitud
	var ciudadUpdate CiudadUpdate

	err := json.NewDecoder(r.Body).Decode(&ciudadUpdate)
	if err != nil {
		handleError(w, "Error al decodificar los datos de actualización", http.StatusBadRequest, err)
		return
	}

	// Actualizar la ciudad en la base de datos según los campos proporcionados
	err = updateCiudad(ciudadUpdate)
	if err != nil {
		handleError(w, "Error al actualizar la ciudad", http.StatusInternalServerError, err)
		return
	}

	// Responder con una respuesta JSON apropiada
	response := map[string]interface{}{
		"status":  "success",
		"message": "Ciudad actualizada con éxito",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// Realizar la actualización en la base de datos para la entidad Ciudad
func updateCiudad(ciudadUpdate CiudadUpdate) error {
	db, err := utils.OpenDB()
	if err != nil {
		return err
	}
	defer db.Close()

	// Construir la consulta SQL basada en los campos no nulos
	query := "UPDATE Ciudad SET"
	args := make([]interface{}, 0)

	if ciudadUpdate.Nombre != nil {
		query += " nombre = $" + strconv.Itoa(len(args)+1) + ","
		args = append(args, *ciudadUpdate.Nombre)
	}

	if ciudadUpdate.PaisID != nil {
		query += " pais_id = $" + strconv.Itoa(len(args)+1) + ","
		args = append(args, *ciudadUpdate.PaisID)
	}

	if len(ciudadUpdate.Imagenes) > 0 {
		// Utiliza pq.Array para convertir las listas de imágenes en un tipo compatible con PostgreSQL
		imagenesArray := pq.Array(ciudadUpdate.Imagenes)
		query += " imagenes = $" + strconv.Itoa(len(args)+1) + ","
		args = append(args, imagenesArray)
	}

	// Eliminar la última coma
	query = query[:len(query)-1]

	// Agregar la condición WHERE
	query += " WHERE id = $" + strconv.Itoa(len(args)+1)
	args = append(args, ciudadUpdate.ID)

	// Ejecutar la consulta
	_, err = db.Exec(query, args...)
	return err
}
