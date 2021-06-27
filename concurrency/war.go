// Learning about concurrency ingo

package main

import (
	"fmt"
	"time"
)

func main() {
	go fight("Defence")
	fight("Attack")
}

func fight(command string) {
	for {
		fmt.Println(command)
		time.Sleep(100 * time.Millisecond)
	}
}
