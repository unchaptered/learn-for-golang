package main

import (
	"fmt"
	"os"
)

func main() {
	addNumber1 := interactive(os.Args[1], os.Args[2])
	addNumber2, mulNumber2 := interactive_function(os.Args[1], os.Args[2])

	fmt.Println("[main]", addNumber1)
	fmt.Println("[main]", addNumber2, mulNumber2)

	username := "Kevin"
	fmt.Println("[main][pointer]", username)

	pointer_function(username)

	fmt.Println("[main][pointer]", username)
}
