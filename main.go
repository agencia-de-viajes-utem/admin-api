package main

import (
	"admin/api/routes"
	"admin/api/utils"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	// Abrir la conexión a la base de datos
	db, err := utils.OpenDB()
	if err != nil {
		log.Fatalf("Error opening database connection: %v", err)
	}
	defer db.Close()

	// Crear un enrutador utilizando gorilla/mux
	r := mux.NewRouter()

	// Configurar rutas desde routes/routes.go
	routes.ConfigureRoutes(r)

	// Imprimir un mensaje si la conexión a la base de datos es exitosa
	log.Println("Connected to the database!")

	// Configurar el servidor web para escuchar en el puerto 9090
	port := os.Getenv("PORT")
	log.Printf("Server is ready! Listening on %s", port)
	log.Fatal(http.ListenAndServe(port, r))
}
