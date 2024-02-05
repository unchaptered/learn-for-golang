package main

import "fmt"

func pointer_function(name string) {
	fmt.Println("[pointer_function]", name)

	name = "John"

	fmt.Println("[pointer_function]", name)
}
