// Learning about goroutines and channels in go

package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan string)
	go count("Moses", channel)

	for msg := range channel {
		fmt.Println(msg)
	}
}

func count(name string, channel chan string) {
	for i := 0; i < 10; i++ {
		channel <- "break"
		fmt.Println(name)
		time.Sleep(time.Millisecond * 100)
	}
	close(channel)
}
