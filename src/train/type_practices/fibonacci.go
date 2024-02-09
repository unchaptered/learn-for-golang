package main

import "fmt"

func fibonacci(n int) []int {
	fmt.Println("[Fibonacci]", n)
	if n < 2 {
		return make([]int, 0)
	}

	nums := make([]int, n)
	nums[0], nums[1] = 1, 1

	for i := 2; i < n; i++ {
		nums[i] = nums[i-2] + nums[i-1]
	}

	return nums
}
