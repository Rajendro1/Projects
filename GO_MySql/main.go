package main

import (
	_ "github.com/go-sql-driver/mysql"
	"main.go/handel"
	"main.go/includes"
)

func main() {
	includes.Connect()
	handel.HandleRequest()
	// r := gin.Default()
	// r.POST("/", func(c *gin.Context) {
	// 	test_image, err := c.FormFile("image")
	// 	if err != nil {
	// 		log.Println(err.Error())
	// 	}
	// 	err1 := c.SaveUploadedFile(test_image, "products_image/"+"try.jpg")
	// 	if err1 != nil {
	// 		log.Println(err1.Error())
	// 	}
	// })
	// r.Run()
}
