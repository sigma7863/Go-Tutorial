package main

import "fmt"

var c, python, java bool // 標準でfalseが入る(bool)

func main() {
	var i int // 標準で0が入る(int)、ローカル(関数内)で定義
	fmt.Println(i, c, python, java) // 0 false false false
}