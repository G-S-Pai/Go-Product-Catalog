// models/product.go
package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	ID       string  `gorm:"type:STRING(36);primaryKey" json:"id"`
	Name     string  `json:"name"`
	Category string  `json:"category"`
	Price    float64 `json:"price"`
}

func (product *Product) BeforeCreate(tx *gorm.DB) (err error) {
    // Generate a new UUID and assign it to the ID field
    product.ID = uuid.New().String()
    return nil
}