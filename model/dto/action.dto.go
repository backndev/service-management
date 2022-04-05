package dto

import "github.com/google/uuid"

type Action struct {
	ID   uuid.UUID `json:"id" gorm:"column:id;primaryKey"`
	Name string    `json:"name" gorm:"column:name"`
}
