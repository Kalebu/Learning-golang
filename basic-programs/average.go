// learning concept of loop  and array in go
// by building a simple program to find average

package main

import "fmt"

func find_average(numbers []int) float32 {
	sum := 0
	for _, no := range numbers {
		sum = sum + no
	}

	average := float32(sum / len(numbers))
	return average
}

func main() {
	var numbers []int
	var n int
	var temp int

	fmt.Println("How many values ? .. ")
	fmt.Scan(&n)
	fmt.Println("Enter ", n, "Values ")
	for i := 0; i < n; i++ {
		fmt.Scan(&temp)
		numbers = append(numbers, temp)
	}

	fmt.Println("The average of ", numbers, " is ", find_average(numbers))
}

// 4
