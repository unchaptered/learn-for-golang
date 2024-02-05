package main

import (
	"fmt"
	"strconv"
)

func interactive(number1 string, number2 string) (result int) {
	int1, _ := strconv.Atoi(number1)
	int2, _ := strconv.Atoi(number2)

	fmt.Println("[interactive]", int1, int2, int1+int2)

	result = int1 + int2
	return
}
