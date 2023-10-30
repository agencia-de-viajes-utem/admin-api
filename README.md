## Descripción

**admin-api** es la esencia misma de la flexibilidad. Un conjunto robusto de rutas y métodos diseñado para brindarte el control total sobre tus operaciones administrativas.

## Rutas y Métodos

- **GET /api/admin/aeropuertos**: Obtén datos sobre aeropuertos.
- **GET /api/admin/ciudades**: Explora información sobre ciudades.
- **GET /api/admin/paquetes**: Descubre detalles sobre paquetes de viaje.
- **GET /api/admin/imagenes/listar**: Explora todas las imágenes disponibles.
- **POST /api/admin/imagenes/subir**: Carga una nueva imagen.
- **GET /api/admin/imagenes/bd**: Accede a imágenes almacenadas en la base de datos.
- **GET /api/admin/imagenes/bucket**: Descubre imágenes almacenadas en el bucket.
- **GET /api/admin/facturacion**: Visualiza todas las facturas.
- **GET /api/admin/facturacion/usuario**: Filtra facturas por usuario.
- **POST /api/admin/facturacion/crear**: Crea una nueva factura.
- **PUT /api/admin/facturacion/actualizar**: Actualiza detalles de una factura.
- **DELETE /api/admin/facturacion/eliminar**: Elimina una factura.
- **GET /api/admin/paises**: Accede a información detallada sobre países.
- **POST /api/admin/paises/crear**: Crea un nuevo país.
- **PUT /api/admin/paises/actualizar**: Actualiza información de un país.

## Configuración

Asegúrate de tener Go instalado en tu máquina.

## Ejecución

1. Clona el repositorio.
2. Ejecuta el comando go run main.go desde la raíz del proyecto.
