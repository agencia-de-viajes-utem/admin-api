package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"cloud.google.com/go/storage"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
)

type SubirImagenResponse struct {
	Mensaje      string `json:"mensaje"`
	RutaEnBucket string `json:"rutaEnBucket"`
}

func PostImagen(w http.ResponseWriter, r *http.Request) {
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("Error loading .env file: %v\n", err)
		http.Error(w, "Error interno del servidor", http.StatusInternalServerError)
		return
	}

	ctx := context.Background()

	// Especifica la ruta a tu archivo JSON de credenciales
	pathToCredentials := "credentials.json"

	// Configura el cliente de Google Cloud Storage con tus credenciales
	client, err := storage.NewClient(ctx, option.WithCredentialsFile(pathToCredentials))
	if err != nil {
		fmt.Printf("Failed to create client: %v\n", err)
		http.Error(w, "Error interno del servidor", http.StatusInternalServerError)
		return
	}

	bucketName := os.Getenv("GCLOUD_BUCKET_NAME")
	carpetaDestino := os.Getenv("CARPETA_DESTINO")

	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	file, header, err := r.FormFile("imagen")
	if err != nil {
		fmt.Printf("Error al leer el archivo: %v\n", err)
		http.Error(w, "Error interno del servidor", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	rutaEnBucket := header.Filename
	if carpetaDestino != "" {
		rutaEnBucket = carpetaDestino + rutaEnBucket
	}
	archivoEnBucket := fmt.Sprintf("gs://%s/%s", bucketName, rutaEnBucket)

	obj := client.Bucket(bucketName).Object(rutaEnBucket)
	_, err = obj.Attrs(ctx)
	if err == nil {
		fmt.Printf("El archivo %s ya existe\n", header.Filename)
		http.Error(w, "Conflicto - El archivo ya existe", http.StatusConflict)
		return
	}

	writer := obj.NewWriter(ctx)
	if _, err := io.Copy(writer, file); err != nil {
		fmt.Printf("Error al subir la imagen: %v\n", err)
		http.Error(w, "Error interno del servidor al subir la imagen", http.StatusInternalServerError)
		return
	}
	if err := writer.Close(); err != nil {
		fmt.Printf("Error al subir la imagen: %v\n", err)
		http.Error(w, "Error interno del servidor al subir la imagen", http.StatusInternalServerError)
		return
	}

	fmt.Printf("Imagen subida con éxito: %s\n", archivoEnBucket)

	response := SubirImagenResponse{
		Mensaje:      "Imagen subida con éxito",
		RutaEnBucket: archivoEnBucket,
	}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		fmt.Printf("Error al serializar la respuesta: %v\n", err)
		http.Error(w, "Error interno del servidor al serializar la respuesta", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)

	// Agregamos logs para imprimir información adicional
	fmt.Printf("Nombre del archivo: %s\n", header.Filename)
	fmt.Printf("Tamaño del archivo: %d bytes\n", header.Size)
	fmt.Printf("Tipo de archivo: %s\n", header.Header.Get("Content-Type"))
	fmt.Printf("Ruta en bucket: %s\n", archivoEnBucket)
}
