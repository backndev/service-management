package dto

import "github.com/google/uuid"

type Role struct {
	ID    uuid.UUID `json:"id"`
	Title string    `json:"title"`
}
