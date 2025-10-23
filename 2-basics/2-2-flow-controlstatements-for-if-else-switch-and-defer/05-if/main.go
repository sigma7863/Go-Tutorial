package main

import (
	"fmt"
	"math"
)

// 虚数の平方根を求める関数
func sqrt(x float64) string {
	// ifもfor同様、ifと{}(中括弧)の間の条件式に()は必要ない
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func main() {
	fmt.Println(sqrt(2), sqrt(-4)) // 1.4142135623730951 2i
}
