package types

import (
	"time"
)

type Movie struct {
	ID        string
	Title     string
	Genre     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
type Movies struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Genre string `json:"genre"`
}
