package main

import (
	"fmt"
	"math"
)

func floating_point_math() {
	const e = 2.71828
	const Avogadro = 6.02214129e23
	const Planck = 6.62606957e-34

	fmt.Println("[floating_point]", math.MaxFloat32, math.MaxFloat64)
	fmt.Println("[floating_point]", e, Avogadro, Planck)
}
