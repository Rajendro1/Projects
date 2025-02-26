package main

import (
	"fmt"
	"net/http"
	"strings"
	"structs"
)

func init() {

}
func main() {
	a := "my f an "
	fmt.Println(strings.TrimSpace(a))
	b := struct {
		Message string
	}{Message: "test"}

}

func GetData(w http.ResponseWriter, r http.Request) {
	a := "my f an "
	fmt.Println(strings.TrimSpace(a))
	b := struct {
		Message string
	}{Message: "test"}
	w.Write(b)
}
