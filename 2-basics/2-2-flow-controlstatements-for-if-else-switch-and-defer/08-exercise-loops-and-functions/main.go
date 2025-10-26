package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	// ↓ ニュートン法を使って求める場合
	z := 1.0
    for i := 0; i < 10; i++ { // 10回ループさせる
        z -= (z*z - x) / (2*z)
    }
    return z

	// ↓ ループ回数ではなく収束条件(差が小さくなったら終わる)で止めるようにする場合
	z := 1.0
	for {
		next := z - (z*z-x)/(2*z)
		if math.Abs(z-next) < 1e-10 { // 誤差が十分小さくなったら終了、math.Abs()は絶対値を返す関数、絶対値を英語で"Absolute value"
			break
		}
		z = next
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2)) // 近似値として求めたもの 1.4142135623746899
	fmt.Println(math.Sqrt(2)) // 実際の値 1.4142135623730951
}
