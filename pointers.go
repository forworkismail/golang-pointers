package main

import "fmt"

func main() {
	age := 32
	name := "test"

	fmt.Println(age)
	fmt.Println(name)

	// Pointer to the age variable
	var myAge *int = &age
	var myName *string = &name
	fmt.Println(myAge)
	fmt.Println(myName)
}
