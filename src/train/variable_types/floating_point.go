package main

import "fmt"

func floating_point() {
	var float_32 float32 = 2_147_483_647
	var float_64 float64 = 9_223_372_036_854_775_807

	fmt.Println("[floating_point]", float_32, float_64)
}
