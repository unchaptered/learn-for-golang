package main

import "fmt"

func basic_list_remove() {

	letters := []string{
		"A",
		"B",
		"C",
		"D",
	}
	remove := 2

	if remove < len(letters) {
		fmt.Println("[basic_list_remove]", letters)
		letters = append(letters[:remove], letters[remove+1:]...)
		fmt.Println("[basic_list_remove]", letters)
	}
}
