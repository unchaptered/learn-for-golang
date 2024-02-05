package main

import "fmt"

func allocate_variable_block() {
	var (
		firstName string = "John"
		lastName string = "Doe"
		age int = 32
	)

	fmt.Println(firstName, lastName, age)
}