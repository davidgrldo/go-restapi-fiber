package models

import "gorm.io/gorm"

// Define a model struct
type Product struct {
	gorm.Model
	Name  string `json:"name"`
	Price uint   `json:"price"`
}
