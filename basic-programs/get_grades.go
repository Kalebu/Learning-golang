// Learning how hash maps works in go

package main

import "fmt"

func main() {
	var subject string
	var database = map[string]int{
		"biology":     100,
		"chemistry":   85,
		"geography":   58,
		"physics":     98,
		"programming": 100,
	}

	for i := 0; i < 10; i++ {
		fmt.Println("Enter subject \t ")
		fmt.Scan(&subject)
		fmt.Println(database[subject])
	}

}
