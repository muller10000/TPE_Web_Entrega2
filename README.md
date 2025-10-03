# Trabajo Práctico 1 / Programación Web

Realizado por **Matias Muller**

## Dominio de la aplicación

El dominio elegido es un **Catálogo de Películas Favoritas**.  
Cada película cuenta con los siguientes atributos:

- **Título**  
- **Descripción**  
- **Estado**: `Vista`, `Pendiente` o `Favorita`  
- **Imagen**: portada o captura representativa de la película  

La página muestra además un resumen estadístico de todas las películas: total, vistas, pendientes y favoritas.

---

## Estructura del proyecto

Peliculas(Entrega1)/
 ├── index.html           # Esqueleto de la página web
 ├── main.go              # Servidor principal en Go
 └── static/
     ├── estilos.css      # Estilos
     └── images/          # Imágenes de cada película en formato JPG

## Requisitos previos

Go
 1.22 o superior

## Ejecución

1) Clonar este repositorio o descargar los archivos.

2) Ejecutar el servidor en Go:
  - cd Peliculas(Entrega1)
  - go run main.go

Abrir en el navegador: http://localhost:8080