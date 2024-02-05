package switch_conditions

import (
	"fmt"
	"math/rand"
	"time"
)

func switch_conditions() {
	secSeed := time.Now().Unix()

	rand.Seed(secSeed)
	randInt := rand.Int31n(10)

	switch randInt {

	case 0:
		fmt.Print("zero...")

	case 1:
		fmt.Print("one...")

	case 2:
		fmt.Print("two...")

	default:
		fmt.Print("no match...")

	}
}
