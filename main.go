package main

import (
	"fmt"
	"github.com/wisphes/filmoteca/internal/handlers"
	"log"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/wisphes/filmoteca/internal/database"
)

func main() {
	db, err := database.NewDB("postgres://user:password@localhost/dbname?sslmode=disable")
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	http.HandleFunc("/actors/add", handlers.AddActorHandler(db))
	http.HandleFunc("/actors/update", handlers.UpdateActorHandler(db))
	http.HandleFunc("/actors/delete", handlers.DeleteActorHandler(db))
	http.HandleFunc("/actors", handlers.GetActorsHandler(db))

	http.HandleFunc("/movies/add", handlers.AddMovieHandler(db))
	http.HandleFunc("/movies/update", handlers.UpdateMovieHandler(db))
	http.HandleFunc("/movies/delete", handlers.DeleteMovieHandler(db))
	http.HandleFunc("/movies", handlers.GetMoviesHandler(db))

	fmt.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
