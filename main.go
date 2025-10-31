package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	_ "github.com/lib/pq"
	"github.com/muller10000/TPE_Web_Entrega2/handlers"
	"github.com/muller10000/TPE_Web_Entrega2/repository"
)

var queries *repository.Queries

// connectDB() devuelve un *sql.DB que se pasa a repository.New(db).

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

func main() {

	// Servir archivos estáticos (de la entrega 1)
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	db, err := connectDB()

	if err != nil {
		panic(fmt.Sprintf("No se pudo conectar a la base de datos: %v", err))
	}

	queries = repository.New(db)

	// Servir index.html en la raíz
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	// GET y POST
	http.HandleFunc("/peliculas", handlers.NewHandlerPeliculas(queries))

	// GET, PUT, DELETE por ID
	http.HandleFunc("/peliculas/", handlers.NewHandlerPeliculaByID(queries))

	fmt.Println("Servidor escuchando en http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error al iniciar servidor:", err)
	}
}
