package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

// Calculator Interface
type calculator interface {
	add() int
	sub() int
	mul() int
	div() (int, error)
}

// Geometric Interface
type geometric interface {
	area() float64
	perim() float64
}

// Calculator Input Struct
type CalculatorInput struct {
	x int
	y int
}

// Geometric Shapes Structs
type triangle struct {
	height float64
	base   float64
	left   float64
	right  float64
}
type square struct {
	width float64
}
type rect struct {
	width, height float64
}
type circle struct {
	radius float64
}

// Calculator Methods
func (i CalculatorInput) add() int {
	return i.x + i.y
}
func (i CalculatorInput) sub() int {
	return i.x - i.y
}
func (i CalculatorInput) mul() int {
	return i.x * i.y
}
func (i CalculatorInput) div() (int, error) {
	if i.y == 0 {
		return 0, errors.New("cannot divide by zero; divisor provided was zero")
	}
	return i.x / i.y, nil
}

// Geometric Methods
func (r rect) area() float64 {
	return r.width * r.height
}
func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}
func (t triangle) area() float64 {
	return t.height * t.base * 0.5
}
func (t triangle) perim() float64 {
	return t.base + t.left + t.right
}
func (s square) area() float64 {
	return s.width * s.width
}
func (s square) perim() float64 {
	return s.width * 4
}

// Main Function to Handle User Input
func main() {
	fmt.Println("Hello! This is your homie calculator.")
	fmt.Println("Please choose what kind of operations you would like to do:")
	fmt.Println("1. Numerical Calculation")
	fmt.Println("2. Geometric Calculation")

	reader := bufio.NewReader(os.Stdin)
	choice, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	choice = strings.TrimSpace(choice)

	switch choice {
	case "1":
		handleCalculator()
	case "2":
		handleGeometricCalculator()
	default:
		fmt.Println("Invalid choice. Please choose 1 or 2.")
	}
}

// Handle Numerical Calculator Operations
func handleCalculator() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the first integer:")
	x, _ := readInt(reader)
	fmt.Println("Enter the second integer:")
	y, _ := readInt(reader)

	calculator := CalculatorInput{x, y}

	fmt.Println("Choose an operation: add, sub, mul, div")
	operation, _ := reader.ReadString('\n')
	operation = strings.TrimSpace(operation)

	switch operation {
	case "add":
		fmt.Printf("Result: %d\n", calculator.add())
	case "sub":
		fmt.Printf("Result: %d\n", calculator.sub())
	case "mul":
		fmt.Printf("Result: %d\n", calculator.mul())
	case "div":
		result, err := calculator.div()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("Result: %d\n", result)
		}
	default:
		fmt.Println("Invalid operation.")
	}
}

// Handle Geometric Calculations
func handleGeometricCalculator() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Choose a shape: square, rectangle, triangle, circle")
	shape, _ := reader.ReadString('\n')
	shape = strings.TrimSpace(shape)

	switch shape {
	case "square":
		fmt.Println("Enter the width of the square:")
		width, _ := readFloat(reader)
		s := square{width}
		printShapeMeasurements(s)
	case "rectangle":
		fmt.Println("Enter the width of the rectangle:")
		width, _ := readFloat(reader)
		fmt.Println("Enter the height of the rectangle:")
		height, _ := readFloat(reader)
		r := rect{width, height}
		printShapeMeasurements(r)
	case "triangle":
		fmt.Println("Enter the base of the triangle:")
		base, _ := readFloat(reader)
		fmt.Println("Enter the height of the triangle:")
		height, _ := readFloat(reader)
		fmt.Println("Enter the left side of the triangle:")
		left, _ := readFloat(reader)
		fmt.Println("Enter the right side of the triangle:")
		right, _ := readFloat(reader)
		t := triangle{height, base, left, right}
		printShapeMeasurements(t)
	case "circle":
		fmt.Println("Enter the radius of the circle:")
		radius, _ := readFloat(reader)
		c := circle{radius}
		printShapeMeasurements(c)
	default:
		fmt.Println("Invalid shape.")
	}
}

// Utility Functions
func printShapeMeasurements(g geometric) {
	fmt.Printf("Area: %.2f\n", g.area())
	fmt.Printf("Perimeter: %.2f\n", g.perim())
}

func readInt(reader *bufio.Reader) (int, error) {
	var input int
	_, err := fmt.Fscanln(reader, &input)
	return input, err
}

func readFloat(reader *bufio.Reader) (float64, error) {
	var input float64
	_, err := fmt.Fscanln(reader, &input)
	return input, err
}
