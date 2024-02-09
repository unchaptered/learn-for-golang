package main

import (
	"fmt"
)

func use_type() {
	var john Employee
	john = Employee{1001, "John", "Doe", "Doe's Street"}

	fmt.Println("[use_type]", john)
}
