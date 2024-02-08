package main

import "fmt"

func recover_function() {
	defer func() {
		handler := recover()

		if handler != nil {
			fmt.Println("[recover]", handler)
		}
	}()

	high_low(10, 1)
	fmt.Println("[recover]")
}
