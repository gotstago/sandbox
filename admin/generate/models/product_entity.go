package models

//see http://blog.ralch.com/tutorial/golang-object-relation-mapping-with-gorm/ for more info on gorm
import "github.com/jinzhu/gorm"

type Product struct {
	gorm.Model
	Description string
	Name        string
}
