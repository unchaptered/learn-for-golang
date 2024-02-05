package main

import "fmt"

func integer_uint_type() {
	var uinteger8 uint8 = 127
	var uinteger16 uint16 = 32_767
	var uinteger32 uint32 = 2_147_483_647
	var uinteger64 uint64 = 9_223_372_036_854_775_807

	fmt.Println("[integer_uint_type]", uinteger8, uinteger16, uinteger32, uinteger64)
}
