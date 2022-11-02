package main

import "fmt"

func main() {

	var months = [...]string{
		"januari",
		"februari",
		"maret",
		"april",
		"mei",
		"juni",
		"juli",
		"agustus",
		"september",
		"oktober",
		"november",
		"desember",
	}

	var slice1 = months[2:5]
	fmt.Println(slice1)

	slice1[0] = "alah sia boy"

	fmt.Println(slice1)
	fmt.Println(months)

	slicebaru := make([]string, 2, 3)

	slicebaru[0] = "hemmm"
	slicebaru[1] = "anwar sahid"
	// slicebaru[2] = "dias"
	// slicebaru[3] = "di"

	slicebaru2 := append(slicebaru, "dias")

	fmt.Println(slicebaru)
	fmt.Println(slicebaru2)
}
