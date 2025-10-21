package main

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i) // complex(複素数型)
)

func main() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)     // Type: bool Value: false
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt) // Type: uint64 Value: 18446744073709551615
	fmt.Printf("Type: %T Value: %v\n", z, z)           // Type: complex128 Value: (2+3i)
}

// Goの型一覧

// bool

// string

// int  int8  int16  int32  int64
// uint uint8 uint16 uint32 uint64 uintptr

// byte // uint8 の別名

// rune // int32 の別名
//      // Unicode のコードポイントを表す

// float32 float64

// complex64 complex128

// (訳注：runeとは古代文字を表す言葉(runes)ですが、Goでは文字そのものを表すためにruneという言葉を使います。)

// 例では、いくつかの型の変数を示しています。また、変数宣言は、インポートステートメントと同様に、まとめて( factored )宣言可能です。

// int, uint, uintptr 型は、32-bitのシステムでは32 bitで、64-bitのシステムでは64 bitです。 サイズ、符号なし( unsigned )整数の型を使うための特別な理由がない限り、整数の変数が必要な場合は int を使うようにしましょう。
