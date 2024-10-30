package main

import "fmt"

func bubbleShort(array []int) []int {
	n := len(array)
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
	return array
}

func main() {
	array := []int{10, 6, 2, 1, 5, 8, 3, 4, 7, 9}
	sorted := bubbleShort(array)
	fmt.Println("Using bubble sort to sort the array", sorted)
}
