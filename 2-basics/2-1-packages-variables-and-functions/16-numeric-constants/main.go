package main

import "fmt"

const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100 // 2¹⁰⁰
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99 // 2¹⁰⁰ ÷ 2⁹⁹ = 2¹ = 2
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	fmt.Println(needInt(Small)) // 21
	fmt.Println(needInt(Big)) // オーバーフロー(int型は64-bitでも9.22×10¹⁸までしか表せない)
	fmt.Println(needFloat(Small)) // 0.2
	fmt.Println(needFloat(Big)) // 1.2676506002282295e+29
}
