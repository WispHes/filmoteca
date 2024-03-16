package models

type Actor struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Gender      string `json:"gender"`
	DateOfBirth string `json:"date_of_birth"`
}
