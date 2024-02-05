package main

import "fmt"

func if_comlex_conditions() {
	num := 5
	if option := true; option {
		fmt.Println("[if_comlex_conditions:if]", option, num)
	} else if num < 10 {
		fmt.Println("[if_comlex_conditions:elseif]", option, num)
	} else {
		fmt.Println("[if_comlex_conditions:else]", option, num)
	}
}
