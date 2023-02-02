package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"main.go/includes"
)

func CreateUser(c *gin.Context) {
	name := c.Request.FormValue("name")
	email := c.Request.FormValue("email")
	password := c.Request.FormValue("password")
	hash, _ := includes.HashPassword(password)
	response, createUser := includes.CreateUserToDb(name, email, hash)
	if createUser == nil {
		c.JSON(http.StatusCreated, gin.H{
			"status":  http.StatusCreated,
			"message": "Your account createing successfully",
			"data":    response,
		})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Message": "Sorry! we can't create your account",
			"status":  http.StatusInternalServerError,
		})
	}
}

func GetAUser(c *gin.Context) {
	userID := c.Request.FormValue("userid")
	getUserUsingUserID, err := includes.GetUsersFrDb(userID)
	if err != nil {
		c.JSON(404, gin.H{
			"status":  404,
			"message": "Sorry!, we don't find your data",
		})
	}
	c.JSON(200, gin.H{
		"status":  200,
		"message": "Your details getting successfully",
		"data":    getUserUsingUserID,
	})
}

func PostLogin(c *gin.Context) {
	email := c.Request.FormValue("email")
	password := c.Request.FormValue("password")
	userByEmail, err := includes.GetPasswordByEmail(email)
	if err != nil {
		c.JSON(404, gin.H{
			"message": "Email id not registered with us, please sign up",
			"status":  404,
		})
	}
	allProductDetails, productErr := includes.GetAllProductsFrDb()

	verifyPassword := includes.CheckPassword(password, userByEmail[0].Password)
	if verifyPassword {
		if productErr != nil || len(allProductDetails) == 0 {
			c.JSON(404, gin.H{
				"message": "No product found",
				"status":  404,
				"data": gin.H{
					"users_data": userByEmail,
				},
			})
		} else {
			c.JSON(200, gin.H{
				"message": "You have successfully logged in",
				"data": gin.H{
					"users_data":   userByEmail,
					"product_data": allProductDetails,
				},
				"status": 200,
			})
		}
	} else {
		c.JSON(401, gin.H{
			"message": "Please give valid password",
			"status":  401,
		})
	}
}

func DeleteAUser(c *gin.Context) {
	userID := c.Request.FormValue("userid")
	getUserUsingUserID, err := includes.GetUsersFrDb(userID)

	if len(getUserUsingUserID) != 0 {
		if includes.DeleteUsersFrDB(userID) {
			c.JSON(200, gin.H{
				"status":  200,
				"message": "Account deleted successfully",
			})
		} else {
			c.JSON(500, gin.H{
				"status":  500,
				"message": "Sorry! we can't delete your account",
			})
		}
	} else if err != nil || len(getUserUsingUserID) == 0 {
		c.JSON(404, gin.H{
			"status":  404,
			"message": "We don't find your account",
		})
	}
}
