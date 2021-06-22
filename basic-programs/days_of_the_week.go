// Learning arrays and loops

package main

import "fmt"

func print_all_days(days [7]string) {
	for index, day := range days {
		fmt.Println(day, " = ", index+1)
	}
}

func number_to_day(days [7]string, number int) {
	fmt.Println(days[number])
}

func main() {
	var days = [...]string{
		"Monday", "Tuesday", "Wednesday", "Thursday",
		"Friday", "Saturday", "Sunday",
	}

	print_all_days(days)
	number_to_day(days, 5)
}
