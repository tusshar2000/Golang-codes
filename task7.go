package main

import "fmt"

type shape interface {
	perimeter() float64
	area() float64
}
type square struct {
	side int
}

func (s *square) perimeter() float64 {
	return float64(4 * s.side)
}
func (s square) area() float64 {
	return float64(s.side * s.side)
}
func printDetails(s shape) {
	fmt.Println("perimeter:", s.perimeter())
	fmt.Println("area:", s.area())

}
func main() {
	var mySquare shape
	mySquare = &square{side: 5}
	printDetails(mySquare)
	fmt.Println(mySquare.perimeter())
	fmt.Println(mySquare.area())
}

// I want you to create an interface named shape
// to implement this interface, a struct needs two methods : perimeter & area
// both should take no arguments and return float64 value
// Create a struct named square which implements shape!
// is this clear?

//those who have completed can try doing the same thing using a function with signature
// func printDetails(s shape) {
// print area
// print perimeter
// }
