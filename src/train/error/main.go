package main

import "fmt"

func main() {
	create_error()
	result1, error1 := create_reusable_error(10)
	result2, error2 := create_reusable_error(1001)

	fmt.Println("[main]", result1, error1)
	fmt.Println("[main]", result2, error2)
}
