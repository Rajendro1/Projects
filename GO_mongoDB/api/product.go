package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"main.go/includes"
)

func CreateProduct(c *gin.Context) {
	productName := c.Request.FormValue("name")
	productPrice := c.Request.FormValue("price")

	product, err := includes.CreateProductToDb(productName, productPrice)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Sorry! we can't add your product",
		})
	}
	c.JSON(http.StatusAccepted, gin.H{
		"status":  http.StatusAccepted,
		"message": "Product creating successfully",
		"data":    product,
	})
}

func GetProduct(c *gin.Context) {
	productID := c.Request.FormValue("id")
	product_details, err := includes.GetProductFrDb(productID)
	if err != nil {
		c.JSON(404, gin.H{
			"status":  404,
			"message": "Sorry!, we don't find your data",
		})
	}
	c.JSON(200, gin.H{
		"status":  200,
		"message": "Product getting successfully",
		"data":    product_details,
	})
}

func DeleteProduct(c *gin.Context) {
	productID := c.Request.FormValue("id")
	product_details, err := includes.GetProductFrDb(productID)
	deleteProduct := includes.DeleteProductFrDB(productID)
	if len(product_details) != 0 {
		if deleteProduct {
			c.JSON(200, gin.H{
				"status":  200,
				"message": "Product deleted successfully",
			})
		} else {
			c.JSON(500, gin.H{
				"status":  500,
				"message": "Sorry! we can't delete your product",
			})
		}
	} else if err != nil || len(product_details) == 0 {
		c.JSON(404, gin.H{
			"status":  404,
			"message": "We don't find your product",
		})
	}
}
func DeleteAllProduct(c *gin.Context) {
	if includes.DeleteAllProductFrDB() {
		c.JSON(200, gin.H{
			"status":  200,
			"message": "All product deleted successfully",
		})
	} else {
		c.JSON(500, gin.H{
			"status":  500,
			"message": "Sorry! we can't deleted all product",
		})
	}
}
func UpdateAllProduct(c *gin.Context) {
	productName := c.Request.FormValue("name")
	productPrice := c.Request.FormValue("price")
	if includes.UpdateAllProductToDB(productName, productPrice) {
		c.JSON(200, gin.H{
			"status":  200,
			"message": "All product updated successfully",
		})
	} else {
		c.JSON(200, gin.H{
			"status":  200,
			"message": "Sorry! we can't update all product",
		})
	}
}
