package handlers

import (
	"context"
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/muller10000/TPE_Web_Entrega2/repository"
)

// define el formato del cuerpo JSON que la API espera recibir

type CreateMovieRequest struct {
	Title    string  `json:"title"`
	Director *string `json:"director"`
	Year     *int32  `json:"year"`
	Genre    *string `json:"genre"`
	Rating   *string `json:"rating"`
}

// Funciones auxiliares

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

// funcion para metodos get y post (CURL)
// Factory para crear el handler pasando queries
// inyeccion de queries: para convertir tu handler en una función que recibe queries como parámetro y así no depender de variables globales

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
