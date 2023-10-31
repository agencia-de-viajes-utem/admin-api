package routes

import (
	aeropuerto "admin/api/handlers/Aeropuertos"
	ciudades "admin/api/handlers/Ciudades"
	facturacion "admin/api/handlers/Facturacion"
	fechapaquete "admin/api/handlers/FechaPaquete"
	home "admin/api/handlers/Home"
	imagenes "admin/api/handlers/Imagenes"
	paises "admin/api/handlers/Paises"
	paquetes "admin/api/handlers/Paquetes"

	"net/http"

	"github.com/gorilla/mux"
)

func ConfigureRoutes(r *mux.Router) {
	// allowedOrigins := []string{"http://admin.lumonidy.studio", "http://localhost:3000"}

	// c := middleware.CorsMiddleware(allowedOrigins)
	// r.Use(c)

	r.Handle("/", http.HandlerFunc(home.HomeHandler)).Methods("GET")
	r.Handle("/check-db", http.HandlerFunc(home.CheckDBHandler)).Methods("GET")
	r.Handle("/static/styles.css", http.HandlerFunc(home.StaticCSSHandler)).Methods("GET")
	r.Handle("/get-routes", http.HandlerFunc(home.GetRoutesHandler)).Methods("GET")

	// Rutas para los Schemas
	r.Handle("/aeropuertos", http.HandlerFunc(aeropuerto.GetAllAeropuertos)).Methods("GET")

	// Ruta para las imágenes
	r.Handle("/imagenes/listar", http.HandlerFunc(imagenes.ListarTodasLasImagenes)).Methods("GET")
	r.Handle("/imagenes/bd", http.HandlerFunc(imagenes.ImagenesBd)).Methods("GET")
	r.Handle("/imagenes/bucket", http.HandlerFunc(imagenes.ImagenesBucket)).Methods("GET")
	r.Handle("/imagenes/subir", http.HandlerFunc(imagenes.PostImagen)).Methods("POST")
	// Facturación
	r.Handle("/facturacion", http.HandlerFunc(facturacion.GetAllFacturas)).Methods("GET")
	r.Handle("/facturacion/usuario", http.HandlerFunc(facturacion.GetFacturasByUsuarios)).Methods("GET")
	r.Handle("/facturacion/crear", http.HandlerFunc(facturacion.CreateFactura)).Methods("POST")
	r.Handle("/facturacion/actualizar", http.HandlerFunc(facturacion.UpdateFactura)).Methods("PUT")
	r.Handle("/facturacion/eliminar", http.HandlerFunc(facturacion.DeleteFactura)).Methods("DELETE")

	// Paises
	r.Handle("/paises", http.HandlerFunc(paises.GetAllPaises)).Methods("GET")
	r.Handle("/paises/crear", http.HandlerFunc(paises.CreatePais)).Methods("POST")
	r.Handle("/paises/actualizar", http.HandlerFunc(paises.UpdatePais)).Methods("PUT")

	//Paquetes
	r.Handle("/paquetes", http.HandlerFunc(paquetes.GetAllPaquetes)).Methods("GET")
	r.Handle("/paquetes/crear", http.HandlerFunc(paquetes.CreatePaquete)).Methods("POST")
	r.Handle("/paquetes/actualizar", http.HandlerFunc(paquetes.UpdatePaquete)).Methods("PUT")
	r.Handle("/paquetes/eliminar", http.HandlerFunc(paquetes.DeletePaquete)).Methods("DELETE")

	//Ciudades
	r.Handle("/ciudades", http.HandlerFunc(ciudades.GetAllCiudades)).Methods("GET")
	r.Handle("/ciudades/crear", http.HandlerFunc(ciudades.CreateCiudad)).Methods("POST")
	r.Handle("/ciudades/actualizar", http.HandlerFunc(ciudades.UpdateCiudad)).Methods("PUT")
	r.Handle("/ciudades/eliminar", http.HandlerFunc(ciudades.DeleteCiudad)).Methods("DELETE")

	//FechaPaquetes
	r.Handle("/fechapaquetes", http.HandlerFunc(fechapaquete.GetAllFechaPaquetes)).Methods("GET")
	r.Handle("/fechapaquetes/mes", http.HandlerFunc(fechapaquete.GetFechaPaqueteByMes)).Methods("GET")
}
