package main

import (
	"fmt"
	"os"
	"strconv"
)

func interactive() {
	number1, _ := strconv.Atoi(os.Args[1])
	number2, _ := strconv.Atoi(os.Args[2])

	fmt.Println("[interactive]", number1, number2, number1+number2)
}
