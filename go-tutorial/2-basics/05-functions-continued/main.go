package main

import "fmt"

// func add(x int, y int) { は以下のように略せる
func add(x, y int) int {
		return x + y // 55
}

func main() {
		fmt.Println(add(42, 13))
}
