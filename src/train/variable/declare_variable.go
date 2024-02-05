package main

import "fmt"

func declare_variable() {
	var firstName, lastName string
	var age int

	firstName = "John"
	lastName = "Doe"
	age = 32

	fmt.Println(firstName, lastName, age)
}