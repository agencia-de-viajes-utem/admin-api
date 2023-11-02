package handlers

import (
	"admin/api/models"
	"encoding/json"
	"log"
	"net/http"
)

func GetAllUsuarios(w http.ResponseWriter, r *http.Request) {
	usuarios, err := fetchUsuarios()
	if err != nil {
		log.Printf("[%d] Error al obtener los usuarios: %v", http.StatusInternalServerError, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	usuariosJSON, err := json.Marshal(usuarios)
	if err != nil {
		log.Printf("[%d] Error al convertir a JSON: %v", http.StatusInternalServerError, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(usuariosJSON)
}

func fetchUsuarios() ([]models.Usuario, error) {
	// Aquí simplemente retornas los datos del mock
	return []models.Usuario{
		{ID: 1, Nombre: "Juan", Apellido: "Pérez", Email: "juan@example.com", Rut: "12345678-9", Fono: "+56912345678", Fotodeperfil: "CrashBiocoot.webp", FechaNacimiento: "1990-05-15"},
		{ID: 2, Nombre: "María", Apellido: "González", Email: "maria@example.com", Rut: "87654321-0", Fono: "+56987654321", Fotodeperfil: "CrashBiocoot.webp", FechaNacimiento: "1985-12-03"},
		{ID: 3, Nombre: "Carlos", Apellido: "Rodríguez", Email: "carlos@example.com", Rut: "11223344-5", Fono: "+56911223344", Fotodeperfil: "CrashBiocoot.webp", FechaNacimiento: "1995-09-20"},
		{ID: 4, Nombre: "Ana", Apellido: "López", Email: "ana@example.com", Rut: "55443322-1", Fono: "+56955443322", Fotodeperfil: "CrashBiocoot.webp", FechaNacimiento: "1988-07-10"},
		{ID: 5, Nombre: "Luis", Apellido: "Martínez", Email: "luis@example.com", Rut: "99887766-4", Fono: "+56999887766", Fotodeperfil: "CrashBiocoot.webp", FechaNacimiento: "1992-04-25"},
	}, nil
}
