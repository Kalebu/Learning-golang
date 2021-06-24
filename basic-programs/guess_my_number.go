// Learning more on std with math

package main

import (
	"fmt"
	"math/rand"
)

func welcome() {
	fmt.Println("---------------------------------")
	fmt.Println("----------WELCOME FRIEND----------")
	fmt.Println("--TESTING YOUR LUCK MADE EASIER---")
}

func propose_a_number(min int, max int) {
	fmt.Println("So here a Luck Test, I have a my secret number")
	fmt.Println("That lies between", min, " and ", max)
}

func get_game_values(min int, max int) int {
	rand_n := rand.Intn(max-min) + min
	return rand_n
}

func check_guess(n int, guess int) bool {
	if n == guess {
		fmt.Println("Congratulation you made !!")
		return true
	} else if n > guess {
		fmt.Println("Guess is too low try again")
	} else {
		fmt.Println("Guess is too high try again")
	}
	return false
}

func main() {
	var secret_number int
	welcome()
	for {
		rand_n := get_game_values(10, 50)
		propose_a_number(10, 50)
		for {
			fmt.Println("Can you guess my number ? ... ")
			fmt.Scan(&secret_number)
			made_t := check_guess(rand_n, secret_number)
			if made_t {
				break
			}
		}

	}
}
