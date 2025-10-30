# Trabajo Práctico Especial - Programación Web

# Autor: Matías Muller

# Proyecto: Películas 3ra Entrega

# Descripción del Proyecto



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



# Requisitos previos

-Linux 

-Go 1.22 o superior

-Docker instalado (Para levantar contenedores)

# Ejecución del proyecto

1) Clonar este repositorio.

2) Copiar el contenido del archivo ".env.example" en un nuevo archivo ".env" reemplazando con las credenciales reales.

DB_NAME=peliculas_tp2
DB_USER=peliculas_user
DB_PASSWORD=peliculas_pass

3) Ejecucion de script en consola linux

./runtest.sh

- Construye la app (build del binario con Docker).
- Levanta los contenedores (DB + API).
- Ejecuta los tests CRUD automáticamente.