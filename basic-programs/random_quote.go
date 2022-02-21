// Learning how to use extenal modules

package main

import (
	"fmt"

	"rsc.io/quote"
)

func main() {
	fmt.Println(quote.Go())
}

// go get rsc.io/quote  (installation)
// still buggy
