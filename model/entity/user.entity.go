package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CheckLogin struct {
	ID       string `json:"id" gorm:"primaryKey, type:varchar(50)"`
	Password string `json:"-" binding:"required"`
}

type User struct {
	gorm.Model
	ID             uuid.UUID `json:"id" gorm:"column:id;primaryKey"`
	PersonalNumber string    `json:"PersonalNumber" gorm:"column:PersonalNumber"`
	Password       string    `json:"password" gorm:"column:password"`
	Email          string    `json:"email" gorm:"column:email"`
	Name           string    `json:"name" gorm:"column:name"`
	RoleId         uuid.UUID `json:"roleId" gorm:"column:roleId"`
	Active         bool      `json:"active" gorm:"column:active"`
}

type UserGetAll struct {
	ID     uuid.UUID `json:"id" gorm:"column:id;primaryKey"`
	Name   string    `json:"name" gorm:"column:name"`
	RoleId uuid.UUID `json:"roleId" gorm:"column:roleId"`
	Title  string    `json:"title" gorm:"column:title"`
	Active bool      `json:"active" gorm:"column:active"`
}

type UserGetById struct {
	ID             uuid.UUID `json:"id" gorm:"column:id;primaryKey"`
	PersonalNumber string    `json:"personalNumber" gorm:"column:personalNumber"`
	Email          string    `json:"email" gorm:"column:email"`
	Name           string    `json:"name" gorm:"column:name"`
	Password       string    `json:"password" gorm:"column:password"`
	RoleId         uuid.UUID `json:"roleId" gorm:"column:roleId"`
	Title          string    `json:"title" gorm:"column:title"`
	Active         bool      `json:"active" gorm:"column:active"`
}
