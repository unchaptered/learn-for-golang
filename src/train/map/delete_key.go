package main

import "fmt"

func delete_key() {
	studentsAge := map[string]int{
		"john": 31,
		"bob":  32,
	}

	fmt.Println("[delete_key]", studentsAge)
	delete(studentsAge, "john")
	fmt.Println("[delete_key]", studentsAge)
}
