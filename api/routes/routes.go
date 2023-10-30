package routes

import (
	"admin/api/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func ConfigureRoutes(r *mux.Router) {
	// allowedOrigins := []string{"http://admin.lumonidy.studio", "http://localhost:3000"}

	// c := middleware.CorsMiddleware(allowedOrigins)
	// r.Use(c)

	r.Handle("/", http.HandlerFunc(handlers.HomeHandler))

	// Rutas para los Schemas
	r.Handle("/aeropuertos", http.HandlerFunc(handlers.GetAllAeropuertos))

	// Ruta para las imágenes
	r.Handle("/imagenes/listar", http.HandlerFunc(handlers.ListarTodasLasImagenes))
	// Ruta para subir imágenes
	r.Handle("/imagenes/subir", http.HandlerFunc(handlers.PostImagen)).Methods("POST")
	// Ruta para ver las imágenes en la bd
	r.Handle("/imagenes/bd", http.HandlerFunc(handlers.ImagenesBd)).Methods("GET")
	// Ruta para ver las imágenes en el bucket
	r.Handle("/imagenes/bucket", http.HandlerFunc(handlers.ImagenesBucket)).Methods("GET")

	// Facturación
	r.Handle("/facturacion", http.HandlerFunc(handlers.GetAllFacturas))
	r.Handle("/facturacion/usuario", http.HandlerFunc(handlers.GetFacturasByUsuarios))
	r.Handle("/facturacion/crear", http.HandlerFunc(handlers.CreateFactura))
	r.Handle("/facturacion/actualizar", http.HandlerFunc(handlers.UpdateFactura))
	r.Handle("/facturacion/eliminar", http.HandlerFunc(handlers.DeleteFactura))

	// Paises
	r.Handle("/paises", http.HandlerFunc(handlers.GetAllPaises))
	r.Handle("/paises/crear", http.HandlerFunc(handlers.CreatePais))
	r.Handle("/paises/actualizar", http.HandlerFunc(handlers.UpdatePais))

	//Paquetes
	r.Handle("/paquetes", http.HandlerFunc(handlers.GetAllPaquetes))
	r.Handle("/paquetes/crear", http.HandlerFunc(handlers.CreatePaquete))
	r.Handle("/paquetes/eliminar", http.HandlerFunc(handlers.DeletePaquete))
	r.Handle("/paquetes/actualizar", http.HandlerFunc(handlers.UpdatePaquete))

	//Ciudades
	r.Handle("/ciudades", http.HandlerFunc(handlers.GetAllCiudades))
	r.Handle("/ciudades/crear", http.HandlerFunc(handlers.CreateCiudad))
	r.Handle("/ciudades/eliminar", http.HandlerFunc(handlers.DeleteCiudad))
	r.Handle("/ciudades/actualizar", http.HandlerFunc(handlers.UpdateCiudad))

	//FechaPaquetes
	r.Handle("/fechapaquetes", http.HandlerFunc(handlers.GetAllFechaPaquetes))
	r.Handle("/fechapaquetes/mes", http.HandlerFunc(handlers.GetFechaPaqueteByMes))
}
