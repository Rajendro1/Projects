package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

func main() {
	fmt.Println(Sum("1.0"))
	eng1 := Engineer[float32]{ID: 123.2}
	printEngineerID(eng1)
}
func Sum[i constraints.Integer | constraints.Float | constraints.Ordered](j i) i {
	return j
}

type EngineerID interface {
	~float32 | ~int32 | ~string
}
type Engineer[E EngineerID] struct {
	ID E
}

func printEngineerID[T EngineerID](e Engineer[T]) {
	fmt.Printf("Engineer ID: %v\n", e.ID)
}
