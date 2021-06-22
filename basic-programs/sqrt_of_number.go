// learning how to use fo std

package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {

	var y float64 = 100.0
	y_sqrt, err := get_sqrt(y)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(y_sqrt)
}

func get_sqrt(number float64) (float64, error) {
	if number < 0 {
		return 0, errors.New("Numbers should be greater than 0 (n>0)")
	}
	return math.Sqrt(float64(number)), nil
}
