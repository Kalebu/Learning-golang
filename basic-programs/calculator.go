package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func minus(x int, y int) int {
	return x - y
}

func multiply(x int, y int) int {
	return x * y
}

func divide(x int, y int) int {
	return x / y
}

func main() {
	var num1 int
	var num2 int
	var oper string

	fmt.Println("Enter a first number:\t")
	fmt.Scan(&num1)
	fmt.Println("Enter Operator:\t")
	fmt.Scan(&oper)
	fmt.Println("Enter a second number:\t")
	fmt.Scan(&num2)

	if oper == "+" {
		answer := add(num1, num2)
		fmt.Println(answer)
	} else if oper == "-" {
		answer := minus(num1, num2)
		fmt.Println(answer)
	} else if oper == "x" {
		fmt.Println(multiply(num1, num2))
	} else if oper == "/" {
		fmt.Println(divide(num1, num2))
	}

}

// 2
