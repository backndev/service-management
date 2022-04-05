package entity

import "github.com/google/uuid"

type Role struct {
	ID     uuid.UUID `json:"id" gorm:"column:id;primaryKey"`
	Title  string    `json:"title" gorm:"column:title"`
	Active bool      `json:"active" gorm:"column:active"`
}
