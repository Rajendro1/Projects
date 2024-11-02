package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

func main() {
	fmt.Println(Sum("1.0"))
}
func Sum[i constraints.Integer | constraints.Float | constraints.Ordered](j i) i {
	return j
}

// type
