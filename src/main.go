package main

import (
	"go-relational-crud/src/domain/models"
	"go-relational-crud/src/main/routes/products"

	"github.com/gin-gonic/gin"
)

func main() {
	p := models.Product{}
	p.Migrate()

	router := gin.Default()

	products.Routes(router)

	router.Run(":8080")
}
