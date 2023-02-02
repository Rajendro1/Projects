package api

import (
	"log"

	"github.com/gin-gonic/gin"
	"main.go/includes"
)

func GetUsers(c *gin.Context) {
	email := c.Request.FormValue("email")
	verifSuperAdminEmail, _ := includes.VerifySuperAdminByEmail(email)
	userDetails, _ := includes.GetAllUsers()
	if verifSuperAdminEmail && len(userDetails) != 0 {
		c.JSON(200, gin.H{
			"message": "Users details getting sucessfully",
			"data":    userDetails,
		})
	} else if !verifSuperAdminEmail {
		c.JSON(401, gin.H{
			"message": "Please give valid super admin email",
		})
	}
}

func PostUsers(c *gin.Context) {
	name := c.Request.FormValue("name")
	email := c.Request.FormValue("email")
	phone := c.Request.FormValue("phone")
	password := c.Request.FormValue("password")
	address := c.Request.FormValue("address")
	mode := c.Request.FormValue("mode")
	hash, _ := includes.HashPassword(password)
	verifyUserEmail, _ := includes.VerifyUsersByEmail(email)

	if includes.IsEmailValid(email) && len(password) >= 8 {
		switch mode {
		case "create_user":
			if !verifyUserEmail {
				createUser, _ := includes.RegisterUser(name, email, phone, address, hash)
				userDetailsByEmail, _ := includes.GetUserDetailsByEmail(email)
				if createUser {
					c.JSON(200, gin.H{
						"message": "User account create successfully",
						"data":    userDetailsByEmail,
					})
				} else {
					c.JSON(500, gin.H{
						"message": "Sorry! we don't create your account",
					})
				}
			} else if verifyUserEmail {
				c.JSON(409, gin.H{
					"message": "This email id already register, please give another one",
				})
			}
		case "login":
			if verifyUserEmail {
				verify_password := includes.CheckPassword(password, includes.GetUsersPassword(email))
				if verify_password {
					userDetailsByEmail, _ := includes.GetUserDetailsByEmail(email)
					productDetailsByUserId, _ := includes.GetAllProductByUsers(int(userDetailsByEmail[0].UserID))
					includes.UpdateUsersLogInTime(int(userDetailsByEmail[0].UserID))
					c.JSON(200, gin.H{
						"message": "Login Successfully",
						"data":    userDetailsByEmail,
						"product": productDetailsByUserId,
					})
				} else if !verify_password {
					c.JSON(401, gin.H{
						"message": "Please give valid password",
					})
				}
			} else if !verifyUserEmail {
				c.JSON(401, gin.H{
					"message": "Email id not register",
				})
			}
		case "logout":

		}
	} else if mode == "logout" {
		log.Println("Log Out Time Start")
		if verifyUserEmail {
			userDetailsByEmail, _ := includes.GetUserDetailsByEmail(email)
			userLogOutTime, _ := includes.UpdateUsersLogOutTime(int(userDetailsByEmail[0].UserID))
			if userLogOutTime {
				c.JSON(200, gin.H{
					"message": "Successfully update the log out time",
				})
			} else if !userLogOutTime {
				c.JSON(500, gin.H{
					"message": "Sorry! we can't update the log out time",
				})
			}
		} else if !verifyUserEmail {
			c.JSON(401, gin.H{
				"message": "Email id not register",
			})
		}
	} else if !includes.IsEmailValid(email) {
		c.JSON(400, gin.H{
			"message": "Please give valid email",
		})
	} else if len(password) < 8 {
		c.JSON(400, gin.H{
			"message": "Please give valid password length minimum 8 characters",
		})
	}
}
func PostSuperAdmin(c *gin.Context) {
	name := c.Request.FormValue("name")
	email := c.Request.FormValue("email")
	phone := c.Request.FormValue("phone")
	password := c.Request.FormValue("password")
	address := c.Request.FormValue("address")
	mode := c.Request.FormValue("mode")
	hash, _ := includes.HashPassword(password)
	verifySuperAdminEmail, _ := includes.VerifySuperAdminByEmail(email)

	if includes.IsEmailValid(email) && len(password) >= 8 {
		switch mode {
		case "create_super_admin":
			if !verifySuperAdminEmail {
				createUser, _ := includes.RegisterSuperAdmin(name, email, phone, address, hash)
				superAdminByEmail, _ := includes.GetSuperAdminByEmail(email)
				if createUser {
					c.JSON(200, gin.H{
						"message": "User account create successfully",
						"data":    superAdminByEmail,
					})
				} else {
					c.JSON(500, gin.H{
						"message": "Sorry! we don't create your account",
					})
				}
			} else if verifySuperAdminEmail {
				c.JSON(409, gin.H{
					"message": "This email id already register, please give another one",
				})
			}
		case "login":
			if verifySuperAdminEmail {
				verify_password := includes.CheckPassword(password, includes.GetSuperAdminPassword(email))
				if verify_password {
					superAdminDetails, _ := includes.GetSuperAdminByEmail(email)
					c.JSON(200, gin.H{
						"message": "Login Successfully",
						"data":    superAdminDetails,
					})
				} else if !verify_password {
					c.JSON(401, gin.H{
						"message": "Please give valid password",
					})
				}
			} else if !verifySuperAdminEmail {
				c.JSON(401, gin.H{
					"message": "Email id not register",
				})
			}
		}
	} else if !includes.IsEmailValid(email) {
		c.JSON(400, gin.H{
			"message": "Please give valid email",
		})
	} else if len(password) < 8 {
		c.JSON(400, gin.H{
			"message": "Please give valid password length minimum 8 characters",
		})
	}
}

func GetUsersActivity(c *gin.Context) {
	email := c.Request.FormValue("email")
	verifyUserEmail, _ := includes.VerifySuperAdminByEmail(email)
	usersActivity, _ := includes.GetUsersActivityBySuperAdmin()

	if verifyUserEmail && includes.IsEmailValid(email) {
		c.JSON(200, gin.H{
			"message": "Users activity getting successfully",
			"data":    usersActivity,
		})
	} else if !includes.IsEmailValid(email) {
		c.JSON(400, gin.H{
			"message": "Please give valid email",
		})
	} else if !verifyUserEmail {
		c.JSON(401, gin.H{
			"message": "Email id not register",
		})
	}
}
