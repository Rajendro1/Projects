package api

import (
	"log"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"main.go/includes"
)

func GetProduct(c *gin.Context) {
	email := c.Request.FormValue("email")
	verifyUserEmail, _ := includes.VerifyUsersByEmail(email)
	userDetailsByEmail, _ := includes.GetUserDetailsByEmail(email)
	productDetailsByUserId, _ := includes.GetAllProductByUsers(int(userDetailsByEmail[0].UserID))
	if verifyUserEmail {
		c.JSON(200, gin.H{
			"message": "Product details are getting successfully",
			"data":    productDetailsByUserId,
		})
	} else if !verifyUserEmail {
		c.JSON(401, gin.H{
			"message": "Email id not register",
		})
	}
}

func PostProduct(c *gin.Context) {
	email := c.Request.FormValue("email")
	// user_id := c.Request.FormValue("user_id")
	name := c.Request.FormValue("name")
	price := c.Request.FormValue("price")
	id := uuid.New()
	image, err := c.FormFile("image")
	if err != nil {
		log.Println(err.Error())
	}

	verifyUserEmail, _ := includes.VerifyUsersByEmail(email)
	if verifyUserEmail {
		userDetailsByEmail, _ := includes.GetUserDetailsByEmail(email)
		str_userID := strconv.Itoa(userDetailsByEmail[0].UserID)
		image_name := str_userID + "+" + id.String() + ".jpg"
		err1 := c.SaveUploadedFile(image, "products_image/"+image_name)
		if err1 != nil {
			log.Println(err1.Error())
		}
		createProduct, _ := includes.CreateProduct(userDetailsByEmail[0].UserID, name, price, image_name)
		if createProduct {
			productDetailsByUserId, _ := includes.GetAllProductByUsers(userDetailsByEmail[0].UserID)
			c.JSON(200, gin.H{
				"message": "Product created successfully",
				"data":    productDetailsByUserId,
			})
		}
	} else if !verifyUserEmail {
		c.JSON(401, gin.H{
			"message": "Email id not register",
		})
	}
}

func UpdateProduct(c *gin.Context) {
	email := c.Request.FormValue("email")
	product_id := c.Request.FormValue("product_id")
	name := c.Request.FormValue("name")
	price := c.Request.FormValue("price")
	image, err := c.FormFile("image")
	if err != nil {
		log.Println(err.Error())
	}
	id := uuid.New()
	int_product_id, _ := strconv.Atoi(product_id)
	verifyUserEmail, _ := includes.VerifyUsersByEmail(email)

	if verifyUserEmail {
		userDetailsByEmail, _ := includes.GetUserDetailsByEmail(email)
		str_userID := strconv.Itoa(userDetailsByEmail[0].UserID)
		image_name := str_userID + "+" + id.String() + ".jpg"
		err1 := c.SaveUploadedFile(image, "products_image/"+image_name)
		if err1 != nil {
			log.Println(err1.Error())
		}
		oldFilePath := includes.SERVER_IMAGE_LOCATION_FOR_DELETE + includes.GetImageNameByProductID(product_id)
		e := os.Remove(oldFilePath)
		if e != nil {
			log.Fatal(e)
		}
		createProduct, _ := includes.UpdateProduct(name, price, image_name, int_product_id, userDetailsByEmail[0].UserID)
		if createProduct {
			productDetailsByUserId, _ := includes.GetAllProductByUsers(userDetailsByEmail[0].UserID)
			c.JSON(200, gin.H{
				"message": "Product update successfully",
				"data":    productDetailsByUserId,
			})
		}
	} else if !verifyUserEmail {
		c.JSON(401, gin.H{
			"message": "Email id not register",
		})
	}
}

func DeleteProduct(c *gin.Context) {
	email := c.Request.FormValue("email")
	product_id := c.Request.FormValue("product_id")
	int_product_id, _ := strconv.Atoi(product_id)
	verifyUserEmail, _ := includes.VerifyUsersByEmail(email)
	if verifyUserEmail {
		userDetailsByEmail, _ := includes.GetUserDetailsByEmail(email)

		oldFilePath := includes.SERVER_IMAGE_LOCATION_FOR_DELETE + includes.GetImageNameByProductID(product_id)
		e := os.Remove(oldFilePath)
		if e != nil {
			log.Fatal(e)
		}

		deleteProduct, _ := includes.DeleteProduct(int_product_id)
		if deleteProduct {
			productDetailsByUserId, _ := includes.GetAllProductByUsers(int(userDetailsByEmail[0].UserID))
			c.JSON(200, gin.H{
				"message": "Product deleted successfully",
				"data":    productDetailsByUserId,
			})
		}
	} else if !verifyUserEmail {
		c.JSON(401, gin.H{
			"message": "Email id not register",
		})
	}
}
