// Learning about achieving concurrency with go select

package main

import (
	"fmt"
	"time"
)

func main() {

	channel_1 := make(chan string)
	channel_2 := make(chan string)

	go func() {
		for {
			channel_1 <- "Every 200ms "
			time.Sleep(time.Millisecond * 200)
		}
	}()

	go func() {
		for {
			channel_2 <- "Every 100ms"
			time.Sleep(time.Millisecond * 100)
		}
	}()

	for {
		select {
		case msg1 := <-channel_1:
			fmt.Println(msg1)
		case msg2 := <-channel_2:
			fmt.Println(msg2)
		}
	}

}
