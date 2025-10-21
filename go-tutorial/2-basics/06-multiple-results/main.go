package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x // world hello
}

func main() {
	a, b := swap("hello", "world") //  :=は短変数宣言(short variable declaration)と言って、var ○○やintなどの型を自動でつけてくれる(関数の中でしか使えない)
	fmt.Println(a, b)
}