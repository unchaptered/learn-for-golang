package main

import "fmt"

func basic_list_bad_copy() {

	letters := []string{
		"A",
		"B",
		"C",
		"D",
		"E",
	}

	slice1 := letters[0:2]
	slice2 := letters[1:4]

	fmt.Println("[basic_list_bad_copy]", "slice1", slice1)
	fmt.Println("[basic_list_bad_copy]", "slice2", slice2)

	slice1[1] = "Z"

	fmt.Println("[basic_list_bad_copy]", "slice1", slice1)
	fmt.Println("[basic_list_bad_copy]", "slice2", slice2)

}
