package database

import (
	"context"
	"github.com/wisphes/filmoteca/internal/models"
)

func (db *DB) AddMovie(ctx context.Context, movie *models.Movie) error {
	_, err := db.ExecContext(ctx, "INSERT INTO movies (title, description, release_date, rating) VALUES ($1, $2, $3, $4)",
		movie.Title, movie.Description, movie.ReleaseDate, movie.Rating)
	if err != nil {
		return err
	}

	if len(movie.Actors) > 0 {
		for _, actorID := range movie.Actors {
			_, err := db.ExecContext(ctx, "INSERT INTO movies_actors (movie_id, actor_id) VALUES ((SELECT id FROM movies ORDER BY id DESC LIMIT 1), $1)", actorID)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (db *DB) UpdateMovie(ctx context.Context, movieID int, movie *models.Movie) error {
	_, err := db.ExecContext(ctx, "UPDATE movies SET title = $1, description = $2, release_date = $3, rating = $4 WHERE id = $5",
		movie.Title, movie.Description, movie.ReleaseDate, movie.Rating, movieID)
	if err != nil {
		return err
	}

	_, err = db.ExecContext(ctx, "DELETE FROM movies_actors WHERE movie_id = $1", movieID)
	if err != nil {
		return err
	}

	if len(movie.Actors) > 0 {
		for _, actorID := range movie.Actors {
			_, err := db.ExecContext(ctx, "INSERT INTO movies_actors (movie_id, actor_id) VALUES ($1, $2)", movieID, actorID)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (db *DB) DeleteMovie(ctx context.Context, movieID int) error {
	_, err := db.ExecContext(ctx, "DELETE FROM movies WHERE id = $1", movieID)
	if err != nil {
		return err
	}

	_, err = db.ExecContext(ctx, "DELETE FROM movies_actors WHERE movie_id = $1", movieID)
	if err != nil {
		return err
	}

	return nil
}

func (db *DB) GetMovies(ctx context.Context) ([]models.Movie, error) {
	rows, err := db.QueryContext(ctx, "SELECT id, title, description, release_date, rating FROM movies")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var movies []models.Movie
	for rows.Next() {
		var movie models.Movie
		if err := rows.Scan(&movie.ID, &movie.Title, &movie.Description, &movie.ReleaseDate, &movie.Rating); err != nil {
			return nil, err
		}

		actors, err := db.getActorsForMovie(ctx, movie.ID)
		if err != nil {
			return nil, err
		}
		movie.Actors = actors

		movies = append(movies, movie)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return movies, nil
}

func (db *DB) getActorsForMovie(ctx context.Context, movieID int) ([]int, error) {
	rows, err := db.QueryContext(ctx, "SELECT actor_id FROM movies_actors WHERE movie_id = $1", movieID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var actorIDs []int
	for rows.Next() {
		var actorID int
		if err := rows.Scan(&actorID); err != nil {
			return nil, err
		}
		actorIDs = append(actorIDs, actorID)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return actorIDs, nil
}
