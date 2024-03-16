package database

import (
	"context"
	"github.com/wisphes/filmoteca/internal/models"
)

func (db *DB) AddActor(ctx context.Context, actor *models.Actor) error {
	_, err := db.ExecContext(ctx, "INSERT INTO actors (name, gender, birthdate) VALUES ($1, $2, $3)",
		actor.Name, actor.Gender, actor.DateOfBirth)
	return err
}

func (db *DB) GetActors(ctx context.Context) ([]models.Actor, error) {
	rows, err := db.QueryContext(ctx, "SELECT id, name, gender, birthdate FROM actors")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var actors []models.Actor
	for rows.Next() {
		var actor models.Actor
		if err := rows.Scan(&actor.ID, &actor.Name, &actor.Gender, &actor.DateOfBirth); err != nil {
			return nil, err
		}
		actors = append(actors, actor)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return actors, nil
}

func (db *DB) UpdateActor(ctx context.Context, actorID int, actor *models.Actor) error {
	_, err := db.ExecContext(ctx, "UPDATE actors SET name = $1, gender = $2, birthdate = $3 WHERE id = $4",
		actor.Name, actor.Gender, actor.DateOfBirth, actorID)
	return err
}

func (db *DB) DeleteActor(ctx context.Context, actorID int) error {
	_, err := db.ExecContext(ctx, "DELETE FROM actors WHERE id = $1", actorID)
	return err
}
