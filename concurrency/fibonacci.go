// Learning about jobs and workers to achieve concurrency
// in go by building program to generate fibonacci numbers

package main

import (
	"fmt"
)

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)
	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)

	for i := 0; i <= 100; i++ {
		jobs <- i
	}

	for result := range results {
		fmt.Println(result)
	}
}

func worker(jobs chan int, results chan int) {
	for n := range jobs {
		results <- fibonacci(n)
	}
	close(jobs)
}

func fibonacci(num int) int {
	if num <= 1 {
		return num
	}
	return fibonacci(num-1) + fibonacci(num-2)
}
