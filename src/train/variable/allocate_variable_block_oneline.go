package main

import "fmt"

func allocate_variable_block_oneline() {
	var (
		firstName, lastName, age = "John", "Doe", 32
	)

	fmt.Println(firstName, lastName, age)
}