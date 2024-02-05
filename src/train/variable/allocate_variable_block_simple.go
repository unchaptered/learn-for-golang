package main

import "fmt"

func allocate_variable_block_simple() {
	var (
		firstName = "John"
		lastName = "Doe"
		age = 32
	)

	fmt.Println(firstName, lastName, age)
}