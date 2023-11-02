// En home.go dentro de la carpeta "handlers"
package handlers

import (
	"admin/api/utils"
	"encoding/json"
	"net/http"
	"path/filepath"
)

// HomeHandler maneja las solicitudes a la ruta "/"
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// Obtener la ruta completa del archivo HTML
	htmlFileName := "index.html"
	htmlFilePath := filepath.Join("api", "utils", htmlFileName)

	// Intentar servir el archivo HTML
	http.ServeFile(w, r, htmlFilePath)
}

func StaticCSSHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "api/utils/styles.css")
}

// GetRoutesHandler maneja las solicitudes para obtener las rutas en formato JSON
func GetRoutesHandler(w http.ResponseWriter, r *http.Request) {
	// Lista de rutas especificadas
	routes := []string{
		"/check-db",
		"/aeropuertos",
		"/imagenes/listar",
		"/imagenes/bd",
		"/imagenes/bucket",
		"/facturacion",
		"/paises",
		"/paquetes",
		"/ciudades",
		"/fechapaquetes",
		"/usuarios",
	}

	// Convertir la lista de rutas a formato JSON
	routesJSON, err := json.Marshal(routes)
	if err != nil {
		http.Error(w, "Error al convertir rutas a JSON", http.StatusInternalServerError)
		return
	}

	// Establecer el encabezado Content-Type y escribir la respuesta
	w.Header().Set("Content-Type", "application/json")
	w.Write(routesJSON)
}

// CheckDBHandler maneja las solicitudes a la ruta "/check-db"
func CheckDBHandler(w http.ResponseWriter, r *http.Request) {
	// Abrir la conexión a la base de datos
	db, err := utils.OpenDB()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}
	defer db.Close()

	// Intentar hacer ping a la base de datos
	err = db.Ping()
	if err != nil {
		// Si hay un error al hacer ping, devolver código 403 (Forbidden)
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("Forbidden"))
		return
	}

	// Si se hizo ping con éxito, devolver código 200 (OK)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
