package main

import "fmt"

func integer_int_type() {
	var integer8 int8 = 127
	var integer16 int16 = -32_767
	var integer32 int32 = 2_147_483_647
	var integer64 int64 = -9_223_372_036_854_775_807

	fmt.Println("[integer_int_type]", integer8, integer16, integer32, integer64)
}
