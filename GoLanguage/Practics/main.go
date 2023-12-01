package main

import (
	// "go/printer"
	"fmt"
)

var path = "app.txt"

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}
	return (err != nil)
}

func main() {
	// fmt.Println("Create DB file")
	// var file, err = os.OpenFile(path, os.O_RDWR, 0644)
	// if isError(err) {
	// 	return
	// }
	// defer file.Close()
	r := []int{1, 3, 1, 2, 5, 7, 1}
	BubbleSort(r)
	Con()
}
func BubbleSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	fmt.Println(arr)
}
func Con() {
	c := make(chan int, 2)
	c <- 1
	a := <-c
	fmt.Println(a)
}
