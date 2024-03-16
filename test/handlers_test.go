package test

import (
	"bytes"
	"encoding/json"
	"github.com/wisphes/filmoteca/internal/database"
	"github.com/wisphes/filmoteca/internal/handlers"
	"github.com/wisphes/filmoteca/internal/models"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestAddActorHandler(t *testing.T) {
	db := database.NewTestDB()
	defer db.Close()

	ts := httptest.NewServer(http.HandlerFunc(handlers.AddActorHandler(db)))
	defer ts.Close()

	actorData := map[string]interface{}{
		"name":        "Test Actor",
		"gender":      "Male",
		"dateOfBirth": "1990-01-01",
	}
	jsonData, err := json.Marshal(actorData)
	if err != nil {
		t.Fatalf("Failed to marshal actor data: %v", err)
	}

	resp, err := http.Post(ts.URL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fatalf("Failed to send POST request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode)
	}
}

func TestGetActorsHandler(t *testing.T) {
	db := database.NewTestDB()
	defer db.Close()

	ts := httptest.NewServer(http.HandlerFunc(handlers.GetActorsHandler(db)))
	defer ts.Close()

	resp, err := http.Get(ts.URL)
	if err != nil {
		t.Fatalf("Failed to send GET request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode)
	}

	var actors []models.Actor
	if err := json.NewDecoder(resp.Body).Decode(&actors); err != nil {
		t.Fatalf("Failed to decode response body: %v", err)
	}

	if len(actors) == 0 {
		t.Errorf("Expected non-empty list of actors")
	}
}

func TestUpdateActorHandler(t *testing.T) {
	db := database.NewTestDB()
	defer db.Close()

	req, err := http.NewRequest("PUT", "/actors/update?id=1", strings.NewReader(`{"name": "Updated Actor", "gender": "Male", "birthdate": "1990-01-01"}`))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(handlers.UpdateActorHandler(db))
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v, want %v", status, http.StatusOK)
	}
}

func TestDeleteActorHandler(t *testing.T) {
	db := database.NewTestDB()
	defer db.Close()

	req, err := http.NewRequest("DELETE", "/actors/delete?id=1", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(handlers.DeleteActorHandler(db))
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v, want %v", status, http.StatusOK)
	}
}

func TestAddMovieHandler(t *testing.T) {
	db := database.NewTestDB()
	defer db.Close()

	req, err := http.NewRequest("POST", "/movies/add", strings.NewReader(`{"title": "Test Movie", "description": "Test Description", "release_date": "2024-03-17", "rating": 8, "actors": [1, 2, 3]}`))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(handlers.AddMovieHandler(db))
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v, want %v", status, http.StatusOK)
	}
}

func TestUpdateMovieHandler(t *testing.T) {
	db := database.NewTestDB()
	defer db.Close()

	req, err := http.NewRequest("PUT", "/movies/update?id=1", strings.NewReader(`{"title": "Updated Movie", "description": "Updated Description", "release_date": "2024-03-17", "rating": 9, "actors": [4, 5, 6]}`))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(handlers.UpdateMovieHandler(db))
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v, want %v", status, http.StatusOK)
	}
}

func TestDeleteMovieHandler(t *testing.T) {
	db := database.NewTestDB()
	defer db.Close()

	req, err := http.NewRequest("DELETE", "/movies/delete?id=1", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(handlers.DeleteMovieHandler(db))
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v, want %v", status, http.StatusOK)
	}
}

func TestGetMoviesHandler(t *testing.T) {
	db := database.NewTestDB()
	defer db.Close()

	req, err := http.NewRequest("GET", "/movies", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(handlers.GetMoviesHandler(db))
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v, want %v", status, http.StatusOK)
	}

	expected := `[{"ID":1,"Title":"Test Movie","Description":"Test Description","ReleaseDate":"2024-03-17","Rating":8,"Actors":null}]`
	if rr.Body.String() != expected {
		t.Errorf("Handler returned unexpected body: got %v, want %v", rr.Body.String(), expected)
	}
}
