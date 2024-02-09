package main

import "fmt"

func main() {
	var num int

	fmt.Println("What's the Fibonacci sequence you want?")
	fmt.Scanln(&num)
	fmt.Println("The Fibonacci sequence is : ", fibonacci(num))

	fmt.Println("MCLX is", romanToArabic("MCLX"))
	fmt.Println("MCMXCIX is ", romanToArabic("MCMXCIX"))
	fmt.Println("MCMZ is", romanToArabic("MCMZ"))
}
