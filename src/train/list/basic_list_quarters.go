package main

import "fmt"

func basic_list_quarters() {
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

	quarters1 := months[0:3]
	fmt.Println("[basic_list_quarters]", "months[0:3]", quarters1)
	fmt.Println("[basic_list_quarters]", "months[0:3]", "len", len(quarters1))
	fmt.Println("[basic_list_quarters]", "months[0:3]", "cap", cap(quarters1))

	quarters2 := months[3:6]
	fmt.Println("[basic_list_quarters]", "months[3:6]", quarters2)
	fmt.Println("[basic_list_quarters]", "months[3:6]", "len", len(quarters2))
	fmt.Println("[basic_list_quarters]", "months[3:6]", "cap", cap(quarters2))

	quarters3 := months[6:9]
	fmt.Println("[basic_list_quarters]", "months[6:9]", quarters3)
	fmt.Println("[basic_list_quarters]", "months[6:9]", "len", len(quarters3))
	fmt.Println("[basic_list_quarters]", "months[6:9]", "cap", cap(quarters3))

	quarters4 := months[9:12]
	fmt.Println("[basic_list_quarters]", "months[9:12]", quarters4)
	fmt.Println("[basic_list_quarters]", "months[9:12]", "len", len(quarters4))
	fmt.Println("[basic_list_quarters]", "months[9:12]", "cap", cap(quarters4))

	quarters5 := months[0:]
	fmt.Println("[basic_list_quarters]", "months[0:]", quarters5)
	fmt.Println("[basic_list_quarters]", "months[0:]", "len", len(quarters5))
	fmt.Println("[basic_list_quarters]", "months[0:]", "cap", cap(quarters5))

	quarters6 := months[6:]
	fmt.Println("[basic_list_quarters]", "months[6:]", quarters6)
	fmt.Println("[basic_list_quarters]", "months[6:]", "len", len(quarters6))
	fmt.Println("[basic_list_quarters]", "months[6:]", "cap", cap(quarters6))

	quarters7 := months[:12]
	fmt.Println("[basic_list_quarters]", "months[:12]", quarters7)
	fmt.Println("[basic_list_quarters]", "months[:12]", "len", len(quarters7))
	fmt.Println("[basic_list_quarters]", "months[:12]", "cap", cap(quarters7))

	quarters8 := months[:6]
	fmt.Println("[basic_list_quarters]", "months[:6]", quarters8)
	fmt.Println("[basic_list_quarters]", "months[:6]", "len", len(quarters8))
	fmt.Println("[basic_list_quarters]", "months[:6]", "cap", cap(quarters8))
}
