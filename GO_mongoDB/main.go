package main

import (
	"main.go/handel"
	"main.go/includes"
)

func main() {
	includes.Connect()
	handel.HandleRequest()
}
