package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim) // ifで定義された変数はelse内でも使用できる
	}
	// can't use v here, though(if, else外では使用できない)
	return lim
}

// ここが理解できなかった...
func main() {
	fmt.Println(
		pow(3, 2, 10), // 27 >= 20
		pow(3, 3, 20), // 9 20
	)
}
