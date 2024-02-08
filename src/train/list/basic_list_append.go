package main

import "fmt"

func basic_list_append() {
	var integers []int

	for i := 0; i < 10; i++ {
		integers = append(integers, i)
		fmt.Println("[basic_list_append]", i, cap(integers), integers)
	}
}
