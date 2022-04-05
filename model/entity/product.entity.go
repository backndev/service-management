package entity

import (
	"github.com/google/uuid"
	"time"
)

type Product struct {
	ID          uuid.UUID `json:"id" gorm:"column:id;primaryKey"`
	Name        string    `json:"name" gorm:"column:name"`
	Description string    `json:"description" gorm:"column:description"`
	Status      string    `json:"status" gorm:"column:status"`
	MakerID     uuid.UUID `json:"makerId" gorm:"column:makerId"`
	SignerID    uuid.UUID `json:"signerId" gorm:"column:signerId"`
	CheckerID   uuid.UUID `json:"checkerId" gorm:"column:checkerId"`
	CreatedAt   time.Time `json:"createdAt" gorm:"column:createdAt"`
	UpdatedAt   time.Time `json:"updatedAt" gorm:"column:updatedAt"`
	DeletedAt   time.Time `json:"deletedAt" gorm:"column:deletedAt"`
}

type ProductGetById struct {
	ID          uuid.UUID `json:"id" gorm:"column:id;primaryKey"`
	Name        string    `json:"name" gorm:"column:name"`
	Description string    `json:"description" gorm:"column:description"`
	Status      string    `json:"status" gorm:"column:status"`
	MakerID     uuid.UUID `json:"makerId" gorm:"column:makerId"`
	SignerID    uuid.UUID `json:"signerId" gorm:"column:signerId"`
	CheckerID   uuid.UUID `json:"checkerId" gorm:"column:checkerId"`
	MakerName   string    `json:"makerName" gorm:"column:makerName"`
	CheckerName string    `json:"checkerName" gorm:"column:checkerName"`
	SignerName  string    `json:"signerName" gorm:"column:signerName"`
}
