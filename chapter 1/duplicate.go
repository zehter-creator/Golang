package main

import "fmt"

//TIP To run your code, right-click the code and select <b>Run</b>. Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.

func main() {
	fmt.Println("duplicate olan her veri ve onların ne kadar tekrarlandığı hakkında")
	myArray := [7]int{1, 2, 2, 3, 3, 3, 4}
	m := make(map[int]int)
	for _, element := range myArray {
		m[element]++
	}
	for key := range m {
		fmt.Println("%d", key)
	}
	for key, value := range m {
		if value > 1 {
			fmt.Println("%d", "%d", key, value)
		}
	}
}
