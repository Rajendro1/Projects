package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	ch := make(chan map[int]bool)

	wg.Add(1)
	go func() {
		for i := 0; i <= 10; i++ {
			Check(i, ch)
		}
	}()
	fmt.Println(<-ch)

	wg.Wait()
}
func Even(n int) {
	if n%2 == 0 {

	}

}
func Odd() {

}
func Check(n int, ch chan map[int]bool) chan map[int]bool {

	if n%2 == 0 {
		ch <- map[int]bool{
			n: true,
		}
		return ch
	} else if n%2 != 0 {
		ch <- map[int]bool{
			n: false,
		}
		return ch
	}

	return ch
}
