package main

import "fmt"

func main() {

	t := triangle{3}
	s := square{4}

	fmt.Println("triangle{3}.perimeter()", t, t.perimeter())
	fmt.Println("square{4}.perimeter()", s, s.perimeter())

	fmt.Println("trinagle{3}.size", t, t.size)
	fmt.Println("square{3}.size", s, s.size)

	t.doubleSize()
	s.doubleSize()

	fmt.Println("trinagle{3}", t, t.size)
	fmt.Println("square{3}", s, s.size)

	originString := originString("Learning Go!")

	fmt.Println(`originString{"Learning Go!"}`, originString)
	fmt.Println(`originString{"Learning Go!"}.Upper()`, originString.Upper())
}
