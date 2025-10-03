# Trabajo Práctico 2 - Programación Web

# Autor: Matías Muller
# Proyecto: Películas TP2

# Descripción del Proyecto

Este proyecto es la segunda entrega de la materia de Programación Web y tiene como objetivo la persistencia de datos de películas en PostgreSQL. Se desarrolló en Go utilizando sqlc para generar código tipado a partir de consultas SQL, evitando errores en tiempo de ejecución. La aplicación permite gestionar películas mediante operaciones CRUD (Crear, Obtener, Listar, Actualizar y Borrar) sobre una tabla movies.

# Dominio de la aplicación

El dominio elegido es Películas.
Cada película cuenta con los siguientes atributos:

id → identificador único

title → título de la película

director → director de la película

year → año de estreno

genre → género de la película

rating → calificación de la película

# Estructura

PeliculasTP2/
 ├── go.mod               # Dependencias del proyecto Go
 ├── go.sum
 ├── index.html           # Página de presentación
 ├── main.go              # Servidor principal / archivo de prueba CRUD
 ├── sqlc.yaml            # Configuración de sqlc para generar código
 ├── database/
 │   ├── queries.sql      # Consultas SQL para sqlc
 │   └── schema.sql       # Definición de la tabla principal
 ├── repository/
 │   ├── db.go            # Configuración de conexión y Queries
 │   ├── models.go        # Definición de modelos en Go
 │   └── queries.sql.go   # Código generado automáticamente por sqlc
 └── static/
     └── estilos.css      # Estilos de la página (opcional)



# Requisitos previos

-Go
 1.22 o superior

-PostgreSQL
 (puede ejecutarse en local o vía Docker)

-sqlc
 instalado en el sistema
(opcional, solo necesario si se desea regenerar el código a partir de los archivos SQL)

-Docker (opcional, recomendado para levantar PostgreSQL)

-Navegador web para visualizar la aplicación

# Ejecución del proyecto

1) Clonar este repositorio o descargar los archivos.

2) Crear la base de datos en PostgreSQL y ejecutar el script database/schema.sql para generar la tabla principal.
(conectarse al puerto 5433 en lugar del puerto por defecto (5432))

3) Opción con Docker
Levantar una instancia de PostgreSQL en el puerto 5433 con:

docker run --name peliculas-db \
  -e POSTGRES_PASSWORD=postgres \
  -p 5433:5432 \
  -d postgres:15

4) Ejecutar en GO
cd Peliculas(Entrega2)
go run main.go

5) Abrir en el navegador:

http://localhost:8080
 → Página de presentación (index.html)

# Tecnologías 
Lenguaje: Go
Base de datos: PostgreSQL
Generador de código: sqlc
Servidor web: Librería estándar net/http de Go
HTML + CSS: para la presentación y los estilos básicos