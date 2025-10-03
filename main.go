package main

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/muller10000/TPE_Web_Entrega2/repository" // remoto
)

func testConnection() {
	// Conexión a la BD
	db, err := sql.Open("postgres", "host=localhost port=5433 user=peliculas_user password=peliculas_pass dbname=peliculas_tp2 sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	queries := repository.New(db)

	// Crear una película de prueba
	movie, err := queries.CreateMovie(context.Background(), repository.CreateMovieParams{
		Title:    "El Padrino",
		Director: sql.NullString{String: "Francis Ford Coppola", Valid: true},
		Year:     sql.NullInt32{Int32: 1972, Valid: true},
		Genre:    sql.NullString{String: "Crimen", Valid: true},
	})

	if err != nil {
		panic(err)
	}
	fmt.Println("Película creada:", movie)

	// Listar todas las películas
	movies, err := queries.ListMovies(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println("Todas las películas:", movies)
}

func main() {
	// Servir archivos estáticos (de la entrega 1)
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Probar la conexión y operaciones CRUD
	testConnection()

	// Handler para la ruta raíz "/"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	fmt.Println("Servidor escuchando en http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error al iniciar servidor:", err)
	}
}
