package main

import "fmt"

func main() {
	nama := "anwar sahid"
	number := 99
	if nama == "anwar" {
		fmt.Println("bener")
	} else if nama == "sahid" {
		fmt.Println("sahid")
	} else {
		fmt.Println("bukan dua dua nya")
	}

	if number%2 == 0 {
		fmt.Println("positif")
	} else {
		fmt.Println("ya negatif")
	}

	if length := len(nama); length < 10 {
		fmt.Println("Oke")
	} else {
		fmt.Println("huruf Kebanyakan")
	}
}
