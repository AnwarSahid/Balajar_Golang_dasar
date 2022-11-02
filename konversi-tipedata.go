package main

import "fmt"

func main() {
	var nilai32 int32 = 30000
	var nilai64 int64 = int64(nilai32)
	var float32 float32 = float32(nilai32)
	fmt.Println(nilai64)
	fmt.Println(float32)

	var name = "Anwar Sahid"
	var getCharacter = name[0]
	var convert = string(getCharacter)

	fmt.Println(convert)

}
