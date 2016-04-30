package models

import "github.com/jinzhu/gorm"

// Create a GORM-backend model
type User struct {
	gorm.Model
	Name string
}
