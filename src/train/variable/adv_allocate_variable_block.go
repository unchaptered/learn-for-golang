package main

import "fmt"

func adv_allocate_variable_block()  {
	firstName, lastName := "John", "Doe"
	age := 32
	
	fmt.Println("[adv_allocate_variable_block]", firstName, lastName, age)
}