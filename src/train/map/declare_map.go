package main

import "fmt"

func declare_map() {
	studentsAge := map[string]int{
		"john": 31,
		"bob":  32,
	}
	fmt.Println("[declare_map]", studentsAge)
}
