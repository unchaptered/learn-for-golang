package main

import "strings"

func (s originString) Upper() string {
	return strings.ToUpper(string(s))
}
