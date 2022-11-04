package main

import "fmt"

func getName() string {
	return "helo word"
}

func getFullname() (string, string) {
	return "anwar", "sahid"
}

func main() {
	namaPertama, namaKedua := getFullname()
	fmt.Println(getName())
	fmt.Println(namaPertama)
	fmt.Println(namaKedua)
}
