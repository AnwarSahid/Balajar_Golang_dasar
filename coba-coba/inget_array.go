package main

import "fmt"

func main() {
	name := [...]string{
		"anwar",
		"sahid",
		"Universitas",
		"Lampung",
		"Lampung Barat",
		"Teknik",
		"Informatika",
	}

	buat_slice := name[2:4]
	// buat_slice[]
	fmt.Println(name)
	buat_slice[0] = "diam"
	fmt.Println(buat_slice)

}
