package main

import "fmt"

func for_loop() {

	sum := 0
	for i := 1; i <= 100; i++ {
		sum += 1
	}

	fmt.Println("[for_loop]", sum)
}
