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

	// to get address of pointer
	fmt.Println(myAge)
	fmt.Println(myName)

	// to get value of pointer
	fmt.Println(*myAge)
	fmt.Println(*myName)

	// to change the value that a pointer is pointing to
	*myAge = 33
	fmt.Println(*myAge) // 33
}
