package handlers

import (
	"context"
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/muller10000/TPE_Web_Entrega3/repository"
)

// define el formato del cuerpo JSON que la API espera recibir

type CreateMovieRequest struct {
	Title    string  `json:"title"`
	Director *string `json:"director"`
	Year     *int32  `json:"year"`
	Genre    *string `json:"genre"`
	Rating   *string `json:"rating"`
}

// Funciones auxiliares para manejar los atributos que pueden ser NULL (ya que estos se convierten en STRUCT)

func valueOrEmpty(s *string) string {
	if s != nil {
		return *s
	}
	return ""
}

func valueOrZero(i *int32) int32 {
	if i != nil {
		return *i
	}
	return 0
}

// Implementación de Handlers CRUD:

// inyeccion de queries: para convertir handler en una función que recibe queries como parámetro y así no depender de variables globales.
// Se realizo "Factory" para crear el handler pasando "queries" como parametro
// Debido a que movi las funciones de main a esta nueva seccion con el fin de modular.

// Crear y listas peliculas

func NewHandlerPeliculas(queries *repository.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			movies, err := queries.ListMovies(context.Background())
			if err != nil {
				panic(err)
			}
			json.NewEncoder(w).Encode(movies)

		case http.MethodPost:
			var req CreateMovieRequest
			if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
				http.Error(w, "JSON inválido", http.StatusBadRequest)
				return
			}

			if req.Title == "" {
				http.Error(w, "El título es obligatorio", http.StatusBadRequest)
				return
			}

			p := repository.CreateMovieParams{
				Title:    req.Title,
				Director: sql.NullString{String: valueOrEmpty(req.Director), Valid: req.Director != nil},
				Year:     sql.NullInt32{Int32: valueOrZero(req.Year), Valid: req.Year != nil},
				Genre:    sql.NullString{String: valueOrEmpty(req.Genre), Valid: req.Genre != nil},
				Rating:   sql.NullString{String: valueOrEmpty(req.Rating), Valid: req.Rating != nil},
			}

			movie, err := queries.CreateMovie(context.Background(), p)
			if err != nil {
				http.Error(w, "Error al crear película", http.StatusInternalServerError)
				return
			}

			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(movie)

		default:
			http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		}
	}
}

// Handler para /peliculas/{id}. Metodos GET, PUT, DELETE por ID

func NewHandlerPeliculaByID(queries *repository.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Extraer el ID desde la URL
		idStr := strings.TrimPrefix(r.URL.Path, "/peliculas/")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "ID inválido", http.StatusBadRequest)
			return
		}

		switch r.Method {
		case http.MethodGet:
			// Obtener película por ID
			movie, err := queries.GetMovie(context.Background(), int32(id))
			if err == sql.ErrNoRows {
				http.Error(w, "Película no encontrada", http.StatusNotFound)
				return
			} else if err != nil {
				http.Error(w, "Error al obtener película", http.StatusInternalServerError)
				return
			}
			json.NewEncoder(w).Encode(movie)

		case http.MethodPut:
			// Actualizar película
			var req CreateMovieRequest
			if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
				http.Error(w, "JSON inválido", http.StatusBadRequest)
				return
			}
			if req.Title == "" {
				http.Error(w, "El título es obligatorio", http.StatusBadRequest)
				return
			}

			params := repository.UpdateMovieParams{
				Title:    req.Title,
				Director: sql.NullString{String: valueOrEmpty(req.Director), Valid: req.Director != nil},
				Year:     sql.NullInt32{Int32: valueOrZero(req.Year), Valid: req.Year != nil},
				Genre:    sql.NullString{String: valueOrEmpty(req.Genre), Valid: req.Genre != nil},
				Rating:   sql.NullString{String: valueOrEmpty(req.Rating), Valid: req.Rating != nil},
				ID:       int32(id),
			}

			movie, err := queries.UpdateMovie(context.Background(), params)
			if err != nil {
				http.Error(w, "Error al actualizar película", http.StatusInternalServerError)
				return
			}
			json.NewEncoder(w).Encode(movie)

		case http.MethodDelete:
			// Verificar si existe
			_, err := queries.GetMovie(context.Background(), int32(id))
			if err == sql.ErrNoRows {
				http.Error(w, "Película no encontrada", http.StatusNotFound)
				return
			} else if err != nil {
				http.Error(w, "Error al buscar película", http.StatusInternalServerError)
				return
			}

			// Eliminar
			if err := queries.DeleteMovie(context.Background(), int32(id)); err != nil {
				http.Error(w, "Error al eliminar película", http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusNoContent)
			w.Write([]byte("Película eliminada correctamente"))

		}
	}
}
