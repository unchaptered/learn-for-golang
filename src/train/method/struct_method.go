package main

func (t triangle) perimeter() int {
	return t.size * 3
}
func (s square) perimeter() int {
	return s.size * 4
}
