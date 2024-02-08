package main

import "fmt"

func high_low(high int, low int) {
	if high < low {
		fmt.Println("[high_low]", "Panic!")
		panic("[high_low][panic] In high_low function, `low` greater than `high`")
	}

	defer fmt.Println("[high_low][defer]", high, low)
	fmt.Println("[high_low][print]", high, low)

	high_low(high, low+1)

}
