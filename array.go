package main

import "fmt"

func main() {
	var name [3]string

	name[0] = "anwar"
	name[1] = "sahid"
	name[2] = "gan"

	fmt.Println(name[0])
	fmt.Println(name[1])
	fmt.Println(name[2])

	var age = [3]int{
		22, 23, 24,
	}

	fmt.Println(age)
	fmt.Println("panjang array", len(age)) /*panjang array bukan banyak data */

}
