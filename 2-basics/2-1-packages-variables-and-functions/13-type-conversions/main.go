package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 3, 4
	// float64(), uint()がないとエラーになる(C言語とは異なり、明示的な変換が必要なため)
	var f float64 = math.Sqrt(float64(x*x + y*y)) // math.Sqrt() : 平方根を計算
	var z uint = uint(f)

	// ↓エラーになる(C言語とは異なり、明示的な変換が必要なため)
	// var f float64 = math.Sqrt((x*x + y*y))
	// var z uint = f

	// ↓でも可
	// f := math.Sqrt(float64(x*x + y*y))
	// z := uint(f)

	fmt.Println(x, y, z) // 3 4 5
}
