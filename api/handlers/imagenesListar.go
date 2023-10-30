package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"cloud.google.com/go/storage"
	"github.com/joho/godotenv"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

// ImagenesBucket obtiene la lista de imágenes en el bucket con el prefijo especificado y las devuelve como JSON.
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

	// Buscar el archivo de credenciales en el directorio actual
	matchingPattern := "./gha-creds-*.json"
	matches, err := filepath.Glob(matchingPattern)
	if err != nil {
		fmt.Println("Error al buscar el archivo de credenciales:", err)
		return nil, err
	}

	if len(matches) == 0 {
		fmt.Println("No se encontraron archivos de credenciales.")
		return nil, fmt.Errorf("no se encontraron archivos de credenciales")
	}

	// Utilizar el primer archivo coincidente (puedes ajustar esto según tus necesidades)
	pathToCredentials := matches[0]

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
