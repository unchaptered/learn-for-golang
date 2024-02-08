package main

import "fmt"

func get_value() {
	studentsAge := make(map[string]int)

	studentsAge["John"] = 32
	studentsAge["Kalha"] = 21

	fmt.Println("[get_value]", studentsAge["Ho"])
	fmt.Println("[get_value]", studentsAge["John"])

	valueHo, hasKeyHo := studentsAge["Ho"]
	fmt.Println("[get_value]", "studentsAge[Ho]", "hasKeyHo", hasKeyHo, "valueHo", valueHo)

	valueJohn, hasKeyJohn := studentsAge["John"]
	fmt.Println("[get_value]", "studentsAge[John]", "hasKeyJohn", hasKeyJohn, "valueJohn", valueJohn)
}
