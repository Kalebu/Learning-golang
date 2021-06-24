// Building more examples with math modules

package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}

func (c circle) get_area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) get_perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func main() {
	duara := circle{radius: 10.43}
	fmt.Println("The area is ", duara.get_area())
	fmt.Println("The Perimeter is ", duara.get_perimeter())
}
