package main

import "fmt"

// SortSlice sorts a slice of integers
func SortSlice(slice []int) {
	for i := 0; i < len(slice); i++ {
		for j := i+1; j < len(slice); j++ {
			if slice[i] > slice[j] {
				slice[i], slice[j] = slice[j], slice[i]
			}
		}
	}
}

// PrintSlice prints given integer slcie
func PrintSlice(slice []int) {
	fmt.Print("[")
	for i, arg := range slice {
		fmt.Print(arg)
		if i != len(slice)-1 {
			fmt.Print(" ")
		}
	}
	fmt.Println("]")
	// OR       fmt.Println(slice)
}

// ReverseSlice reverses slice
func ReverseSlice(slice []int) {
	for i, j := len(slice)-1, 0; i > j; i, j = i-1, j+1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
}

// IcrementOdd increments odd-positioned numbers in integer slice
func IncrementOdd(slice []int) {
	for i := 0; i < len(slice); i++ {
		if i % 2 == 1 {
			slice[i]++
		}
	}
}

// appendFunc appends functions and returns a new function
func appendFunc(dst func([]int), src ...func([]int)) func([]int) {
	return func(slice []int) {
		dst(slice)
		for _, f := range src {
			f(slice)
		}
	}
}
