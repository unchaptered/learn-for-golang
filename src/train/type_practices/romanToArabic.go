package main

import "fmt"

func romanToArabic(roman string) int {
	romanMapper := map[rune]int{
		'M': 1000,
		'D': 500,
		'C': 100,
		'L': 50,
		'X': 10,
		'V': 5,
		'I': 1,
	}

	arabicVals := make([]int, len(roman)+1)
	for idx, digit := range roman {
		if val, isPresent := romanMapper[digit]; isPresent {
			arabicVals[idx] = val
		} else {
			fmt.Printf("Error: The roman numeral %s has a bad digit: %c\n", roman, digit)
			return 0
		}
	}

	arabic := 0

	for idx := 0; idx < len(roman); idx++ {
		if arabicVals[idx] < arabicVals[idx+1] {
			arabicVals[idx] = -arabicVals[idx]
		}

		arabic += arabicVals[idx]
	}

	return arabic
}
