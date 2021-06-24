// Learning about go runtime module

package main

import (
	"fmt"
	"runtime"
)

func main() {

	os := runtime.GOOS

	switch os {
	case "linux":
		fmt.Println("For this you need skills ")
	case "windows":
		fmt.Println("For this you need patient")
	default:
		fmt.Println("Well for not yet , looks cool")
	}
}
