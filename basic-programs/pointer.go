// The basics of pointer in go

package main

import "fmt"

func main() {
	x := 50
	add(&x, 12)
	add(&x, 40)
	add(&x, 90)
	fmt.Println(x)
}

func add(number *int, value int) {
	*number = *number + value
}
