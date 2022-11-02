package main

import "fmt"

func main() {

	name := "anwar sahid"
	leng := len(name)

	switch {
	case leng < 10:
		fmt.Println("sudah oke")
	case leng > 10:
		fmt.Println("Namenya Kebanyakan")
	}
}
