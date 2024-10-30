package main

import "fmt"

func mergeShort(array []int) []int {
	if len(array) < 2 {
		return array
	}
	mid := len(array) / 2
	left := mergeShort(array[:mid])
	right := mergeShort(array[mid:])

	return merge(left, right)
}
func merge(left, right []int) []int {
	var result []int
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)
	return result
}

func main() {
	array := []int{10, 6, 2, 1, 5, 8, 3, 4, 7, 9}
	sorted := mergeShort(array)
	fmt.Println("Sorted by using mergeShort method: ", sorted)
}
