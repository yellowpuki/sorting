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

func findMax(array []int) int {
	max, maxIdx := 0, 0

	for i := range array {
		if array[i] > max {
			max, maxIdx = array[i], i
		}
	}
	return maxIdx
}

func insSort(array []int) {
	var x int
	for i := 1; i < len(array); i++ {
		x = array[i]
		j := i

		for ; j > 0 && array[j-1] > x; j-- {
			array[j] = array[j-1]
		}

		array[j] = x
	}
}

func main() {
	var a = []int{3, 1, 5, 7, 8, 2, 6}
	fmt.Println("Min el index ", findMin(a), " value = ", a[findMin(a)])
	fmt.Println("Max el index ", findMax(a), " value = ", a[findMax(a)])

	insSort(a)

	fmt.Println("Sorted array: ", a)
}
