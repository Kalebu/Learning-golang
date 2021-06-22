// Learning concept of loop in go

package main

import "fmt"

func main() {
	var limit int
	var number int

	fmt.Println("Generate Multiplication Table")
	fmt.Println("Enter Target number ")
	fmt.Scan(&number)
	fmt.Println("Enter limit of the table ")
	fmt.Scan(&limit)

	for i := 1; i <= limit; i++ {
		fmt.Println(number, "x", i, "=", i*number)
	}
}

// 3
