package main

import "fmt"

func split(sum int) (x, y int) { // (x, y int)を付けることでreturnに何も書かなくて済む(長い関数では可読性が下がるので付けない)
	x = sum * 4 / 9
	y = sum - x
	return // 7 10
}

func main() {
	fmt.Println(split(17)) // func split()を呼び出し
}