package main

import "fmt"

func loop_of_map() {
	studentsAge := map[string]int{
		"john": 31,
		"bob":  32,
	}

	for name, age := range studentsAge {
		fmt.Printf("[loop_of_map] %s\t%d\n", name, age)
	}
}
