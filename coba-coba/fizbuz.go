package main

import "fmt"

func main() {
	value := 15

	if value%3 == 0 && value%5 == 0 {
		fmt.Println("fizBuz")
	} else if value%3 == 0 {
		fmt.Println("fiz")
	} else if value%5 == 0 {
		fmt.Println("buzz")
	} else {
		fmt.Println(value)
	}
}
