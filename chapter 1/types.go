package chapter_1

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perimeter() float64
}
type rectangle struct {
	width, height float64
}
type circle struct {
	radius float64
}

func (r rectangle) area() float64 {
	return r.width * r.height
}
func (r rectangle) perimeter() float64 {
	return 2*r.width + 2*r.height
}
func (r circle) area() float64 {
	return r.radius * r.radius * math.Pi
}
func (r circle) perimeter() float64 {
	return 2*r.radius + 2*math.Pi
}
func measure(g geometry) float64 {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perimeter())
	return math.Sqrt(g.perimeter())
}
func main() {
	r := rectangle{width: 10, height: 5}
	c := circle{radius: 5}
	measure(r)
	measure(c)
}
