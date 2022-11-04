package main

import "fmt"

func main() {

	// counter := 1

	// for counter <= 10 {
	// 	fmt.Println("perulangan", counter)

	// 	counter++
	// }

	// for i := 0; i < 10; i++ {
	// 	fmt.Println("perulangan", i)
	// }

	slice := []string{
		"anwar",
		"sahid",
		"sada",
	}

	// for i := 0; i < len(slice); i++ {
	// 	fmt.Println(slice[i])
	// }

	for _, data_slice := range slice {
		fmt.Println(data_slice)
	}

}
