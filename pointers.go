package main

import "fmt"

func main() {
	age := 32

	fmt.Println(age)

	// Pointer to the age variable
	var myAge *int = &age
	fmt.Println(myAge)
}
