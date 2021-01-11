package models

import (
	"go-relational-crud/src/infra/db/sqlite"

	"gorm.io/gorm"
)

var db = sqlite.Connect()

// Product is model
type Product struct {
	gorm.Model
	Name  string
	Code  string
	Price float64
}

// Migrate is a function to migrate Products
func (p Product) Migrate() {
	db.AutoMigrate(p)
}

// FindAll is a method which returns all Products
func (p Product) FindAll(product *[]Product) {
	db.Find(&product)
}

// Create is a function to Create Products
func (p Product) Create(product *Product) {
	db.Create(&product)
}

// Update a product based on id
func (p Product) Update(product *Product, id uint64) {
	var prdt *Product
	db.Model(&prdt).Where("id = ?", id).Updates(&product)
}

// Delete is a method to delete a Product
func (p Product) Delete(id uint64) {
	db.Delete(&Product{}, id)
}
