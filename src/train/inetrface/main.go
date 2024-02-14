package main

import "fmt"

func main() {
	var s Shape = Square{3}

	fmt.Println("Square{3}.Area()", s.Area())
	fmt.Println("Square{3}.Perimeter()", s.Perimeter())
}
