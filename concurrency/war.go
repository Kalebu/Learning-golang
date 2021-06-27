// Learning about concurrency ingo

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		fight("Attack")
		wg.Done()
	}()

	go func() {
		fight("Defence")
		wg.Done()
	}()

	wg.Wait()
}

func fight(command string) {
	for i := 0; i <= 10; i++ {
		fmt.Println(command)
		time.Sleep(100 * time.Millisecond)
	}
}
