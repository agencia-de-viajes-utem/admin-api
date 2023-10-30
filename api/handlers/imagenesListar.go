package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"cloud.google.com/go/storage"
	"github.com/joho/godotenv"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

// ListarImagenes obtiene la lista de imágenes en el bucket con el prefijo especificado y las devuelve como JSON.
func ImagenesBucket(w http.ResponseWriter, r *http.Request) {
	infoImagenes, err := fetchInfoImagenes()
	if err != nil {
		handleError(w, "Error al obtener la lista de imágenes", http.StatusInternalServerError, err)
		return
	}

	infoImagenesJSON, err := json.Marshal(infoImagenes)
	if err != nil {
		handleError(w, "Error al convertir a JSON", http.StatusInternalServerError, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(infoImagenesJSON)
}

func fetchInfoImagenes() ([]map[string]string, error) {
	// Cargar variables de entorno desde el archivo .env
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error al cargar las variables de entorno:", err)
		return nil, err
	}

	// Especificar la ruta al archivo JSON de credenciales
	pathToCredentials := "./credentials.json"

	// Configurar el cliente de Google Cloud Storage con las credenciales
	ctx := context.Background()
	client, err := storage.NewClient(ctx, option.WithCredentialsFile(pathToCredentials))
	if err != nil {
		fmt.Println("Error al configurar el cliente de Google Cloud Storage:", err)
		return nil, err
	}

	defer client.Close()

	bucketName := os.Getenv("GCLOUD_BUCKET_NAME")
	carpetaDestino := os.Getenv("CARPETA_DESTINO")

	// Obtener la lista de archivos en el bucket con el prefijo especificado
	it := client.Bucket(bucketName).Objects(ctx, &storage.Query{Prefix: carpetaDestino})
	var infoImagenes []map[string]string

	for {
		archivo, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			fmt.Println("Error al iterar sobre los archivos:", err)
			return nil, err
		}

		// Filtrar directorios
		if !strings.HasSuffix(archivo.Name, "/") {
			infoImagen := map[string]string{
				"nombre":      archivo.Name[len(carpetaDestino):], // Eliminar el prefijo
				"url_publica": fmt.Sprintf("https://storage.googleapis.com/%s/%s", bucketName, archivo.Name),
			}
			infoImagenes = append(infoImagenes, infoImagen)
		}
	}

	return infoImagenes, nil
}
