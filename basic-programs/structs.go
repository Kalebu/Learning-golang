// Learning how to create structs in go

package main

import "fmt"

type Anonymous struct {
	name       string
	age        int
	motivation string
}

func main() {
	elliot := Anonymous{name: "Elliot", age: 12, motivation: "Human rights"}
	fmt.Println("Famous Anonymous  member details")
	fmt.Println("Name:  ", elliot.name)
	fmt.Println("Age: ", elliot.age)
	fmt.Println("Motivation for hacking: ", elliot.motivation)
}
