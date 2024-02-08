package main

import "fmt"

func make_map() {
	studentsAge := make(map[string]int)

	studentsAge["John"] = 32
	studentsAge["Kalha"] = 21

	fmt.Println("[make_map]", studentsAge)
}
