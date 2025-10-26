package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	for i := 0; i < 20; i++ {
		next := z - (z*z-x)/(2*z)
		fmt.Printf("Step %2d: z = %.15f (差=%.15f)\n", i, z, math.Abs(z-next))
		if math.Abs(z-next) < 1e-10 {
			fmt.Println("→ 収束しました！")
			break
		}
		z = next
	}
	return z
}

func main() {
	fmt.Printf("自作Sqrt(2) = %.15f\n", Sqrt(2))
	fmt.Printf("math.Sqrt(2) = %.15f\n", math.Sqrt(2))
	// Step  0: z = 1.000000000000000 (差=0.500000000000000)
	// Step  1: z = 1.500000000000000 (差=0.083333333333333)
	// Step  2: z = 1.416666666666667 (差=0.002450980392157)
	// Step  3: z = 1.414215686274510 (差=0.000002123899820)
	// Step  4: z = 1.414213562374690 (差=0.000000000001595)
	// → 収束しました！
	// 自作Sqrt(2) = 1.414213562374690
	// math.Sqrt(2) = 1.414213562373095
}
