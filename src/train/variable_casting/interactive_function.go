package main

import (
	"fmt"
	"strconv"
)

func interactive_function(number1 string, number2 string) (result_a int, result_b int) {
	int1, _ := strconv.Atoi(number1)
	int2, _ := strconv.Atoi(number2)

	fmt.Println("[interactive]", int1, int2, int1+int2)

	result_a = int1 + int2
	result_b = int1 * int2
	return
}
