package products

import (
	"go-relational-crud/src/domain/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Routes is a method to shelter group requests
func Routes(router *gin.Engine) {
	p := router.Group("/products")
	{
		p.GET("/", func(context *gin.Context) {
			var products []models.Product
			prdt := models.Product{}
			prdt.FindAll(&products)

			context.JSON(http.StatusOK, products)
		})
		p.POST("/", func(context *gin.Context) {
			name := context.PostForm("name")
			code := context.PostForm("code")
			price, err := strconv.ParseFloat(context.PostForm("price"), 32)

			if err != nil {
				panic("Impossible to convert.")
			}

			p := models.Product{Name: name, Code: code, Price: price}

			// prdt := models.Product{}
			p.Create(&p)
			context.JSON(http.StatusOK, &p)
		})
		p.PATCH("/:id", func(context *gin.Context) {
			name := context.PostForm("name")
			code := context.PostForm("code")
			price, err := strconv.ParseFloat(context.PostForm("price"), 32)

			if err != nil {
				panic("Impossible to convert string to float")
			}

			id, err := strconv.ParseUint(context.Param("id"), 5, 64)

			if err != nil {
				panic("Impossible to convert string to int")
			}

			p := models.Product{Name: name, Code: code, Price: price}
			p.Update(&p, id)

			context.JSON(http.StatusOK, p)
		})
		p.DELETE("/:id", func(context *gin.Context) {
			id, err := strconv.ParseUint(context.Param("id"), 5, 64)

			if err != nil {
				panic("Impossible to convert string to int.")
			}

			p := models.Product{}
			p.Delete(id)

			context.JSON(http.StatusOK, gin.H{
				"message": "The product with id: " + strconv.FormatUint(id, 10) + " was deleted.",
			})
		})
	}
}
