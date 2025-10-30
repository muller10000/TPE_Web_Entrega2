package main

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"time"

	_ "github.com/lib/pq"
	"github.com/muller10000/TPE_Web_Entrega2/handlers"
	"github.com/muller10000/TPE_Web_Entrega2/repository"
)

var queries *repository.Queries

func connectDB() (*sql.DB, error) {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	return sql.Open("postgres", connStr)
}

func testConnection() {
	var db *sql.DB
	var err error

	// Intentar conectarse con reintentos (la BD puede demorar en levantarse)
	for i := 1; i <= 5; i++ {
		db, err = connectDB()
		if err == nil {
			err = db.Ping()
			if err == nil {
				break
			}
		}
		fmt.Printf("Esperando base de datos (%d/5)...\n", i)
		time.Sleep(3 * time.Second)
	}
	if err != nil {
		panic(fmt.Sprintf("No se pudo conectar a la base de datos: %v", err))
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
	fmt.Println("✅ Película creada:", movie)

	// Listar todas las películas
	movies, err := queries.ListMovies(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println("🎬 Todas las películas:", movies)
}

func main() {
	// Servir archivos estáticos (de la entrega 1)
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Probar conexión y CRUD
	testConnection()

	db, err := connectDB()

	if err != nil {
		panic(fmt.Sprintf("No se pudo conectar a la base de datos: %v", err))
	}

	queries = repository.New(db)

	// Handler principal
	//http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//	http.ServeFile(w, r, "index.html")
	//})

	http.HandleFunc("/peliculas", handlers.NewHandlerPeliculas(queries))

	fmt.Println("Servidor escuchando en http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error al iniciar servidor:", err)
	}
}
