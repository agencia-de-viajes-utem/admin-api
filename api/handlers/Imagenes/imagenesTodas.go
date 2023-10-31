package handlers

import (
	"admin/api/models"
	"encoding/json"
	"log"
	"net/http"
)

// ListarTodasLasImagenes combina las imágenes de la base de datos y del bucket y las devuelve como JSON.
func ListarTodasLasImagenes(w http.ResponseWriter, r *http.Request) {
	// Obtener imágenes de la base de datos
	imagenesBd, err := FetchBdInfoImagenes()
	if err != nil {
		log.Printf("[%d] Error al obtener las imágenes de la base de datos: %v", http.StatusInternalServerError, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Obtener imágenes del bucket
	imagenesBucket, err := fetchInfoImagenes()
	if err != nil {
		log.Printf("[%d] Error tratando de obtener las imágenes desde el bucket: %v", http.StatusInternalServerError, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Combinar imágenes por ID/nombre
	imagenesCombinadas := combinarImagenes(imagenesBd, imagenesBucket)

	// Convertir a JSON y enviar como respuesta
	imagenesJSON, err := json.Marshal(imagenesCombinadas)
	if err != nil {
		log.Printf("[%d] Error al convertir a JSON: %v", http.StatusInternalServerError, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(imagenesJSON)
}

// combinarImagenes combina las imágenes de la base de datos y del bucket por ID/nombre.
func combinarImagenes(imagenesBd []models.Imagen, imagenesBucket []map[string]string) []map[string]string {
	imagenesCombinadas := make([]map[string]string, 0)

	// Crear un mapa para indexar imágenes de la base de datos por ID
	imagenesBdMap := make(map[string]models.Imagen)
	for _, imagen := range imagenesBd {
		imagenesBdMap[imagen.ID] = imagen
	}

	// Combinar imágenes
	for _, imagenBucket := range imagenesBucket {
		id := imagenBucket["nombre"]
		imagenBd, ok := imagenesBdMap[id]

		// Si la imagen existe en la base de datos
		if ok {
			imagenCombinada := map[string]string{
				"id":             imagenBd.ID,
				"alt":            imagenBd.Alt,
				"descripcion":    imagenBd.Descripcion,
				"fecha_creacion": imagenBd.FechaCreacion.Format("2006-01-02 15:04:05"),
				"url_publica":    imagenBucket["url_publica"],
			}
			imagenesCombinadas = append(imagenesCombinadas, imagenCombinada)
		} else {
			// Si la imagen no existe en la base de datos, agregar con campos vacíos
			imagenCombinada := map[string]string{
				"id":             id,
				"alt":            "",
				"descripcion":    "",
				"fecha_creacion": "",
				"url_publica":    imagenBucket["url_publica"],
			}
			imagenesCombinadas = append(imagenesCombinadas, imagenCombinada)
		}
	}

	return imagenesCombinadas
}
