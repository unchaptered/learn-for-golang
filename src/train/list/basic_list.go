package main

import "fmt"

func basic_list() {
	months := []string{
		"January",
		"February",
		"March",
		"April",
		"May",
		"June",
		"July",
		"August",
		"September",
		"October",
		"November",
		"December",
	}
	fmt.Println("[basic_list]", months)
	fmt.Println("[basic_list] len", len(months))
	fmt.Println("[basic_list] cap", cap(months))
}
