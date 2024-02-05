package main

import "fmt"

func declare_variable_block() {
	var (
		firstName, lastName string
		age int
	)

	firstName = "John"
	lastName = "Doe"
	age = 32

	fmt.Println(firstName, lastName, age)
}