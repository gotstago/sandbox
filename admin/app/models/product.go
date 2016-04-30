package models

//see http://blog.ralch.com/tutorial/golang-object-relation-mapping-with-gorm/ for more info on gorm
import "github.com/jinzhu/gorm"

// Product creates another GORM-backend model
type Product struct {
	gorm.Model
	Name        string
	Description string
}