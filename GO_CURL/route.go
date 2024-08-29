package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func HandleRoute() {
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {

		singleLineCommand := c.Request.FormValue("data")
		singleLineCommand = strings.ReplaceAll(singleLineCommand, "\n", " ")
		singleLineCommand = strings.ReplaceAll(singleLineCommand, "\t", "") // Optional: remove tabs if necessary
		singleLineCommand = strings.ReplaceAll(singleLineCommand, `curl `, "")
		singleLineCommand = strings.ReplaceAll(singleLineCommand, `    `, "")
		singleLineCommand = strings.ReplaceAll(singleLineCommand, `\\`, "")
		singleLineCommand = strings.ReplaceAll(singleLineCommand, `\`, "")
		c.JSON(http.StatusOK, singleLineCommand)
		parsed, err := ParseCurlCommand(singleLineCommand)
		if err != nil {
			fmt.Println("Error parsing cURL command:", err)
			return
		}

		response, err := ExecuteCurlCommand(parsed)
		if err != nil {
			fmt.Println("Error executing cURL command:", err)
			return
		}
		c.JSON(http.StatusOK, gin.H{"response": response})

	})
	r.GET("/new", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "DON")
	})

	r.Run(":80")
}
