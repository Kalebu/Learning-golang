// Learning how to use time module in go

package main

import (
	"fmt"
	"time"
)

func factorial(n float64) float64 {
	fact := 1
	for i := 1; i <= int(n); i++ {
		fact *= i
	}
	return float64(fact)
}

func main() {

	start := time.Now()
	var number float64
	fmt.Println("Factorial calculator !!")
	fmt.Println("Enter a any number ")
	fmt.Scan(&number)
	result := factorial(number)
	fmt.Println("The factorial of ", number, " is ", result)
	now := time.Now()
	time_taken := now.Sub(start)
	fmt.Println("The time taken is ", time_taken, " Seconds ")
}
