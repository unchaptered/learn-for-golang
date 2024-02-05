package main

import (
	"fmt"
	"math/rand"
	"time"
)

func pre_and_post_statement() {
	var num int64

	rand.Seed(time.Now().Unix())
	for num != 5 {
		num = rand.Int63n(15)
		fmt.Println("[pre_and_post_statement]", num)
	}
}
