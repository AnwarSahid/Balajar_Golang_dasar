package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		if i == 9 {
			break
		} else if i%3 == 0 {
			continue
		}

		fmt.Println(i)
	}
}
