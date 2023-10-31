package handlers

import (
	"admin/api/utils"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/lib/pq" // Importa la librería pq para trabajar con arrays PostgreSQL
)

// CiudadCreate representa la estructura para la creación de una nueva ciudad
type CiudadCreate struct {
	Nombre   string   `json:"nombre"`
	PaisID   int      `json:"pais_id"`
	Imagenes []string `json:"imagenes"`
	// Otras propiedades de ciudad, si las hubiera
}

func CreateCiudad(w http.ResponseWriter, r *http.Request) {
	var ciudadCreate CiudadCreate

	if err := json.NewDecoder(r.Body).Decode(&ciudadCreate); err != nil {
		log.Printf("[%d] Error al decodificar los datos de creación: %v", http.StatusBadRequest, err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := createCiudad(ciudadCreate); err != nil {
		log.Printf("[%d] Error al crear la ciudad: %v", http.StatusInternalServerError, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"status":  "success",
		"message": "Ciudad creada con éxito",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func createCiudad(ciudad CiudadCreate) error {
	// Realizar la verificación en la base de datos
	db, err := utils.OpenDB()
	if err != nil {
		return err
	}
	defer db.Close()

	// Verificar si ya existe una ciudad con el mismo nombre en el mismo país
	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM ciudad WHERE nombre = $1 AND pais_id = $2", ciudad.Nombre, ciudad.PaisID).Scan(&count)
	if err != nil {
		return err
	}

	// Si count es mayor que cero, significa que ya existe una ciudad con el mismo nombre en el mismo país
	if count > 0 {
		return fmt.Errorf("ya existe una ciudad con el mismo nombre en el mismo país")
	}

	// Utiliza pq.Array para convertir las listas de imágenes en un tipo compatible con PostgreSQL
	imagenesArray := pq.Array(ciudad.Imagenes)

	// Realizar la inserción en la base de datos
	_, err = db.Exec(`
        INSERT INTO ciudad (nombre, pais_id, imagenes)
        VALUES ($1, $2, $3)
    `, ciudad.Nombre, ciudad.PaisID, imagenesArray)

	return err
}
