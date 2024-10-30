package main

import (
	"fmt"
	"sort"
)

// Ortalama (Mean) hesaplama
func calculateMean(array []float64) float64 {
	sum := 0.0
	for _, num := range array {
		sum += num
	}
	return sum / float64(len(array))
}

// Ortanca (Median) hesaplama
func calculateMedian(array []float64) float64 {
	sort.Float64s(array) // Ortanca hesaplamak için diziyi sıralıyoruz
	n := len(array)
	if n%2 == 0 {
		return (array[n/2-1] + array[n/2]) / 2 // Çift sayıda eleman varsa ortadaki iki sayının ortalamasını alır
	}
	return array[n/2] // Tek sayıda eleman varsa ortadaki eleman ortancadır
}

// Mod (Mode) hesaplama
func calculateMode(array []float64) float64 {
	counts := make(map[float64]int)
	for _, num := range array {
		counts[num]++
	}

	maxCount := 0
	mode := array[0]
	for num, count := range counts {
		if count > maxCount {
			maxCount = count
			mode = num
		}
	}
	return mode
}

func main() {
	array := []float64{1, 2, 2, 3, 4, 4, 4, 5, 5}

	mean := calculateMean(array)
	median := calculateMedian(array)
	mode := calculateMode(array)

	fmt.Printf("Mean (Ortalama): %.2f\n", mean)
	fmt.Printf("Median (Ortanca): %.2f\n", median)
	fmt.Printf("Mode (Mod): %.2f\n", mode)
}
