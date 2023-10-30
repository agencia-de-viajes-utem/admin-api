package handlers

import (
	"admin/api/utils"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/lib/pq" // Importa la librería pq para trabajar con arrays PostgreSQL
)

// PaqueteCreate representa la estructura para la creación de un nuevo paquete
type PaqueteCreate struct {
	Nombre      string   `json:"nombre"`
	IdOrigen    int      `json:"id_origen"`
	IdDestino   int      `json:"id_destino"`
	Descripcion string   `json:"descripcion"`
	Detalles    string   `json:"detalles"`
	PrecioVuelo float64  `json:"precio_vuelo"`
	IdHh        []int    `json:"id_hh"`
	Imagenes    []string `json:"imagenes"`
}

func CreatePaquete(w http.ResponseWriter, r *http.Request) {
	// Obtener los datos para la creación del paquete desde el cuerpo de la solicitud
	var paqueteCreate PaqueteCreate

	err := json.NewDecoder(r.Body).Decode(&paqueteCreate)
	if err != nil {
		handleError(w, "Error al decodificar los datos de creación", http.StatusBadRequest, err)
		return
	}

	// Crear el nuevo paquete en la base de datos
	err = createPaquete(paqueteCreate)
	if err != nil {
		handleError(w, "Error al crear el paquete", http.StatusInternalServerError, err)
		return
	}

	// Responder con éxito y proporcionar una respuesta JSON
	response := map[string]interface{}{
		"status":  "success",
		"message": "Paquete creado con éxito",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func createPaquete(paquete PaqueteCreate) error {
	// Realizar la verificación en la base de datos
	db, err := utils.OpenDB()
	if err != nil {
		return err
	}
	defer db.Close()

	// Verificar si ya existe un paquete con el mismo nombre (u otro criterio que desees)
	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM paquete WHERE nombre = $1", paquete.Nombre).Scan(&count)
	if err != nil {
		return err
	}

	// Si count es mayor que cero, significa que ya existe un paquete con el mismo nombre
	if count > 0 {
		return fmt.Errorf("ya existe un paquete con el mismo nombre")
	}

	// Utiliza pq.Array para convertir las listas en un tipo compatible con PostgreSQL
	idHhArray := pq.Array(paquete.IdHh)
	imagenesArray := pq.Array(paquete.Imagenes)

	// Realizar la inserción en la base de datos
	_, err = db.Exec(`
    INSERT INTO paquete (nombre, id_origen, id_destino, descripcion, detalles, precio_vuelo, id_hh, imagenes)
    VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
`, paquete.Nombre, paquete.IdOrigen, paquete.IdDestino, paquete.Descripcion, paquete.Detalles, paquete.PrecioVuelo, idHhArray, imagenesArray)

	return err
}
