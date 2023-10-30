package handlers

import (
	"admin/api/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

// PaisCreate representa la estructura para la creación de un nuevo país
type PaisCreate struct {
	Nombre      string `json:"nombre"`
	Abreviacion string `json:"abreviacion"`
	// Otras propiedades de país, si las hubiera
}

func CreatePais(w http.ResponseWriter, r *http.Request) {
	// Obtener los datos para la creación del país desde el cuerpo de la solicitud
	var paisCreate PaisCreate

	err := json.NewDecoder(r.Body).Decode(&paisCreate)
	if err != nil {
		handleError(w, "Error al decodificar los datos de creación", http.StatusBadRequest, err)
		return
	}

	// Crear el nuevo país en la base de datos
	err = createPais(paisCreate.Nombre, paisCreate.Abreviacion)
	if err != nil {
		handleError(w, "Error al crear el país", http.StatusInternalServerError, err)
		return
	}

	// Responder con éxito y proporcionar una respuesta JSON
	response := map[string]interface{}{
		"status":  "success",
		"message": "País creado con éxito",
	}

	// Puedes incluir información adicional, como el país recién creado, si lo deseas
	// response["country"] = paisCreate

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func createPais(nombre string, abreviacion string) error {
	// Realizar la verificación en la base de datos
	db, err := utils.OpenDB()
	if err != nil {
		return err
	}
	defer db.Close()

	// Verificar si ya existe un país con el mismo nombre o abreviación
	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM pais WHERE nombre = $1 OR abreviacion = $2", nombre, abreviacion).Scan(&count)
	if err != nil {
		return err
	}

	// Si count es mayor que cero, significa que ya existe un país con el mismo nombre o abreviación
	if count > 0 {
		return fmt.Errorf("ya existe un país con el mismo nombre o abreviación")
	}

	// Realizar la inserción en la base de datos
	_, err = db.Exec(`
        INSERT INTO pais (nombre, abreviacion)
        VALUES ($1, $2)
    `, nombre, abreviacion)

	return err
}
