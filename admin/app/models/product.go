package models

import "github.com/jinzhu/gorm"

// Product creates another GORM-backend model
type Product struct {
	gorm.Model
	Name        string
	Description string
}