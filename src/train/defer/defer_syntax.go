package main

import "fmt"

func defer_syntax() {
	for i := 1; i <= 4; i++ {
		defer fmt.Println("[defer_syntax][defer]", -i)
		fmt.Println("[defer_syntax]", i)
	}
}
