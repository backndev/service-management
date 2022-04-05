package dto

import (
	"github.com/google/uuid"
)

type User struct {
	ID             uuid.UUID `json:"id" gorm:"column:id;primaryKey"`
	PersonalNumber string    `json:"PersonalNumber" gorm:"column:PersonalNumber"`
	Password       string    `json:"password" gorm:"column:password"`
	Email          string    `json:"email" gorm:"column:email"`
	Name           string    `json:"name" gorm:"column:name"`
	Role           Role      `json:"role" gorm:"foreignKey:RoleId"`
	Active         bool      `json:"active" gorm:"column:active"`
}

type UserGetAll struct {
	ID     uuid.UUID `json:"id" gorm:"column:id;primaryKey"`
	Name   string    `json:"name" gorm:"column:name"`
	Role   Role      `json:"role"`
	Active bool      `json:"active" gorm:"column:active"`
}

type UserGetById struct {
	ID             uuid.UUID `json:"id" gorm:"column:id;primaryKey"`
	PersonalNumber string    `json:"personalNumber" gorm:"column:personalNumber"`
	Email          string    `json:"email" gorm:"column:email"`
	Name           string    `json:"name" gorm:"column:name"`
	Role           Role      `json:"role" gorm:"foreignKey:RoleId"`
	Active         bool      `json:"active" gorm:"column:active"`
}

type UserLogin struct {
	PersonalNumber string `json:"PersonalNumber" binding:"required"`
	Password       string `json:"password" binding:"required"`
}
