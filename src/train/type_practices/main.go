package main

import "fmt"

func main() {
	var num int

	fmt.Println("What's the Fibonacci sequence you want?")
	fmt.Scanln(&num)
	fmt.Println("The Fibonacci sequence is : ", fibonacci(num))
}
