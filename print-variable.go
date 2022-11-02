package main

import "fmt"

func main() {
	var nama string
	var universitas = "Uniersitas Lampung"
	age := 22

	// bisa juga pake variable multiple

	var (
		firstname = "sahid"
		lastname  = "sahid"
	)

	nama = "anwar sahid"
	fmt.Println(nama, universitas, age, firstname)

	nama = "sahid anwar"
	fmt.Println(nama, universitas, age, lastname)
}
