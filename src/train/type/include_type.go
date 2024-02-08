package main

type Person struct {
	ID        int
	FirstName string
	LastName  string
	Address   string
}

type Employee struct {
	Information Person
	ManagerID   int
}

func use_type() {
}
