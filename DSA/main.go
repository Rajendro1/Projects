package main

import (
	"fmt"
)

func main() {
	// str := "fmttmf"
	// GetReverseString(str)
	// a := []int{1, -1, 4, 5, 2, 4, 12, 999, 7, 8, 10, 11}
	// FindTheTargetElement(a, 11)
	// FindTheMaxAndMinValueFromTheSlice(a)
	// fmt.Println(FectorialNumber(5))
	// fmt.Println(customSorting([]string{"5",
	// 	"abc",
	// 	"ab",
	// 	"abcde",
	// 	"t",
	// 	"a",
	// 	"abcd"}))
	// var a int
	// fmt.Scan(&a)
	// fmt.Println(a)
	// miniMaxSum([]int32{7, 69, 2, 221, 8974})
	a := rotateLeft(int32(4), []int32{1, 2, 3, 4, 5})
	fmt.Println(a)
}

// 1,2,3,4,5
// 2,3,4,5,1
func rotateLeft(k int32, a []int32) []int32 {
	var out []int32
	for i := 0; i < int(k); i++ {
		for j := 1; j < len(a); j++ {
			out = append(out, a[j])
		}
		out = append(out, a[0])

		a = out
		out = []int32{}
	}
	return a
}
func rotateLeftV2(k int32, a []int32) []int32 {
	n := int32(len(a))
	k = k % n // Reduce k to a value less than the length of the array

	// Create a new slice to store the rotated array
	rotated := make([]int32, n)

	// Calculate the effective rotation index for each element
	for i := int32(0); i < n; i++ {
		rotatedIndex := (i - k + n) % n
		rotated[rotatedIndex] = a[i]
	}

	return rotated
}

func circularArrayRotation(a []int32, k int32, queries []int32) []int32 {
	var out []int32
	var final []int32
	for i := 0; i < int(k); i++ {
		out = append(out, a[len(a)-1])
		for j := 0; j < len(a)-1; j++ {
			out = append(out, a[j])
		}
		a = out
		out = []int32{}
	}
	for m := 0; m < len(queries); m++ {
		final = append(final, a[queries[m]])
	}
	return final
}
func circularArrayRotationV2(a []int32, k int32, queries []int32) []int32 {
	n := int32(len(a))
	k = k % n // Reduce k to a value less than the length of the array

	var final []int32

	// Calculate the effective rotation index for each query and append the corresponding value to final
	for _, q := range queries {
		// Calculate the effective index after rotation
		rotatedIndex := (q - k + n) % n
		final = append(final, a[rotatedIndex])
	}

	return final
}

func miniMaxSum(arr []int32) {
	// Write your code here
	var sum1, sum2, total int32
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {

			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
		total += arr[i]
		fmt.Println(total)
	}
	sum1 = total - arr[len(arr)-1]
	sum2 = total - arr[0]
	fmt.Println(total, arr[0], arr[len(arr)-1])
	fmt.Println(arr)
	fmt.Println(sum1, sum2)

}

func customSorting(strArr []string) []string {
	newSlice := []string{}
	for i := 0; i < len(strArr); i++ {
		found := false
		for j := i + 1; j < len(strArr); j++ {
			if strArr[i] == strArr[j] {
				found = true
				break

			}
		}
		if !found {
			newSlice = append(newSlice, strArr[i])
		}
	}
	return newSlice
}
func GetReverseString(s string) {
	d := len(s) / 2
	firstStrSlice := []string{}
	secondStrSlice := []string{}
	for i := 0; i < d; i++ {
		firstStrSlice = append(firstStrSlice, string(s[i]))
	}

	for i := len(s) - 1; i >= d; i-- {
		secondStrSlice = append(secondStrSlice, string(s[i]))
	}

	for i := 0; i < len(firstStrSlice); i++ {
		if firstStrSlice[i] != secondStrSlice[i] {
			fmt.Println("false")
			return
		}
	}
	fmt.Println("true")
}
func GetTargetNumberFromSlice() {
	a := []int{3, 2, 1, 4, 5}
	target := 10
	var t1, t2, t3 int
	for i := 0; i < len(a)-2; i++ { // Fix: len(a)-2 instead of len(a)
		for j := i + 1; j < len(a)-1; j++ { // Fix: i+1 instead of i
			for k := j + 1; k < len(a); k++ {
				if a[i]+a[j]+a[k] == target {
					t1 = a[i]
					t2 = a[j]
					t3 = a[k]
				}
			}
		}
	}
	fmt.Println(t1, t2, t3)
}
func FindTheTargetElement(n []int, target int) {
	for i := 0; i < len(n); i++ {
		if n[i] == target {
			fmt.Println(n[i])
		}
	}
}
func FindTheMaxAndMinValueFromTheSlice(a []int) {
	for j, _ := range a {
		for i := j; i < len(a); i++ {
			if a[j] > a[i] {
				a[j], a[i] = a[i], a[j]
			}
		}
	}
	fmt.Println(a)
	fmt.Println(a[len(a)-1])
	fmt.Println(a[0])
}

func FectorialNumber(num int) int {
	if num == 0 {
		return 1
	}
	return num * FectorialNumber(num-1)
}
