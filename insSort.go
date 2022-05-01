package main

import (
	"fmt"
	"math"
)

func findMin(array []int) int {
	min := math.MaxInt
	minIdx := 0

	for i := range array {
		if array[i] < min {
			min, minIdx = array[i], i
		}
	}
	return minIdx
}

func main() {
	var a = []int{3, 1, 5, 7, 8, 2, 6}
	fmt.Println("Min el nimber ", findMin(a), " value = ", a[findMin(a)])
}
