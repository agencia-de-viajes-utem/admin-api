package handlers

import (
	"admin/api/utils"
	"encoding/json"
	"net/http"
	"strconv"
)

// PaisUpdate representa la estructura para la actualización parcial de un país
type PaisUpdate struct {
	ID          int     `json:"id"`
	Nombre      *string `json:"nombre,omitempty"`
	Abreviacion *string `json:"abreviacion,omitempty"`
	Imagenes    *string `json:"imagenes,omitempty"`
}

// Handler que llama a la función updatePais y devuelve una respuesta JSON
func UpdatePais(w http.ResponseWriter, r *http.Request) {
	// Obtener los datos actualizados del país desde el cuerpo de la solicitud
	var paisUpdate PaisUpdate

	err := json.NewDecoder(r.Body).Decode(&paisUpdate)
	if err != nil {
		handleError(w, "Error al decodificar los datos de actualización", http.StatusBadRequest, err)
		return
	}

	// Actualizar el país en la base de datos según los campos proporcionados
	err = updatePais(paisUpdate)
	if err != nil {
		handleError(w, "Error al actualizar el país", http.StatusInternalServerError, err)
		return
	}

	// Responder con una respuesta JSON apropiada
	response := map[string]interface{}{
		"status":  "success",
		"message": "Pais actualizado con éxito",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// Realizar la actualización en la base de datos
func updatePais(paisUpdate PaisUpdate) error {

	db, err := utils.OpenDB()
	if err != nil {
		return err
	}
	defer db.Close()

	// Construir la consulta SQL basada en los campos no nulos
	query := "UPDATE Pais SET"
	args := make([]interface{}, 0)

	if paisUpdate.Nombre != nil {
		query += " nombre = $" + strconv.Itoa(len(args)+1) + ","
		args = append(args, *paisUpdate.Nombre)
	}

	if paisUpdate.Abreviacion != nil {
		query += " abreviacion = $" + strconv.Itoa(len(args)+1) + ","
		args = append(args, *paisUpdate.Abreviacion)
	}

	if paisUpdate.Imagenes != nil {
		query += " imagenes = $" + strconv.Itoa(len(args)+1) + ","
		args = append(args, *paisUpdate.Imagenes)
	}

	// Eliminar la última coma
	query = query[:len(query)-1]

	// Agregar la condición WHERE
	query += " WHERE id = $" + strconv.Itoa(len(args)+1)
	args = append(args, paisUpdate.ID)

	// Ejecutar la consulta
	_, err = db.Exec(query, args...)
	return err
}
