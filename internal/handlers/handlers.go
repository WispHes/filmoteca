package handlers

import (
	"encoding/json"
	"github.com/wisphes/filmoteca/internal/database"
	"github.com/wisphes/filmoteca/internal/models"
	"net/http"
	"strconv"
)

func AddActorHandler(db *database.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var actor models.Actor
		err := json.NewDecoder(r.Body).Decode(&actor)
		if err != nil {
			http.Error(w, "Failed to decode request body", http.StatusBadRequest)
			return
		}

		err = db.AddActor(r.Context(), &actor)
		if err != nil {
			http.Error(w, "Failed to add actor to database", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}

func DeleteActorHandler(db *database.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		actorIDStr := r.URL.Query().Get("id")
		if actorIDStr == "" {
			http.Error(w, "Actor ID is required", http.StatusBadRequest)
			return
		}

		actorID, err := strconv.Atoi(actorIDStr)
		if err != nil {
			http.Error(w, "Invalid actor ID", http.StatusBadRequest)
			return
		}

		err = db.DeleteActor(r.Context(), actorID)
		if err != nil {
			http.Error(w, "Failed to delete actor from database", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func UpdateActorHandler(db *database.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		actorIDStr := r.URL.Query().Get("id")
		if actorIDStr == "" {
			http.Error(w, "Actor ID is required", http.StatusBadRequest)
			return
		}

		actorID, err := strconv.Atoi(actorIDStr)
		if err != nil {
			http.Error(w, "Invalid actor ID", http.StatusBadRequest)
			return
		}

		var updatedActor models.Actor
		err = json.NewDecoder(r.Body).Decode(&updatedActor)
		if err != nil {
			http.Error(w, "Failed to decode request body", http.StatusBadRequest)
			return
		}

		err = db.UpdateActor(r.Context(), actorID, &updatedActor)
		if err != nil {
			http.Error(w, "Failed to update actor in database", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func GetActorsHandler(db *database.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		actors, err := db.GetActors(r.Context())
		if err != nil {
			http.Error(w, "Failed to fetch actors from database", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(actors); err != nil {
			http.Error(w, "Failed to encode actors into JSON", http.StatusInternalServerError)
			return
		}
	}
}

func AddMovieHandler(db *database.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var movie models.Movie
		err := json.NewDecoder(r.Body).Decode(&movie)
		if err != nil {
			http.Error(w, "Failed to decode request body", http.StatusBadRequest)
			return
		}

		err = db.AddMovie(r.Context(), &movie)
		if err != nil {
			http.Error(w, "Failed to add movie to database", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}

func UpdateMovieHandler(db *database.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		movieIDStr := r.URL.Query().Get("id")
		if movieIDStr == "" {
			http.Error(w, "Movie ID is required", http.StatusBadRequest)
			return
		}

		movieID, err := strconv.Atoi(movieIDStr)
		if err != nil {
			http.Error(w, "Invalid movie ID", http.StatusBadRequest)
			return
		}

		var updatedMovie models.Movie
		err = json.NewDecoder(r.Body).Decode(&updatedMovie)
		if err != nil {
			http.Error(w, "Failed to decode request body", http.StatusBadRequest)
			return
		}

		err = db.UpdateMovie(r.Context(), movieID, &updatedMovie)
		if err != nil {
			http.Error(w, "Failed to update movie in database", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func DeleteMovieHandler(db *database.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		movieIDStr := r.URL.Query().Get("id")
		if movieIDStr == "" {
			http.Error(w, "Movie ID is required", http.StatusBadRequest)
			return
		}

		movieID, err := strconv.Atoi(movieIDStr)
		if err != nil {
			http.Error(w, "Invalid movie ID", http.StatusBadRequest)
			return
		}

		err = db.DeleteMovie(r.Context(), movieID)
		if err != nil {
			http.Error(w, "Failed to delete movie from database", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func GetMoviesHandler(db *database.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		movies, err := db.GetMovies(r.Context())
		if err != nil {
			http.Error(w, "Failed to fetch movies from database", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(movies); err != nil {
			http.Error(w, "Failed to encode movies into JSON", http.StatusInternalServerError)
			return
		}
	}
}
