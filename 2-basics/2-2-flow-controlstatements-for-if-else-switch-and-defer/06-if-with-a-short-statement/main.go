package main

import (
	"fmt"
	"math"
)

// pow()はべき乗を求める関数, 型はfloat64でなければならない | powはpowerの略(べき乗という意味がある)
func pow(x, n, lim float64) float64 { // x: もとになる数, n: 指数(何回かけるか). lim: 上限値
	if v := math.Pow(x, n); v < lim { // vはifの中でのみ有効
		return v
	}
	return lim // ここでvを返そうとするとvはスコープの外なのでエラーになる
}

func main() {
	fmt.Println(
		pow(3, 2, 10), // 9
		pow(3, 3, 20), // 20(3^3は27だが、limが20のため、20と出力している)
	)
}
