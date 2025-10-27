package main

import "fmt"

func main() {
	// deferは後か最後に実行したい(呼び出し元の関数の終わり(returnする))ときに使う
	defer fmt.Println("world")

	fmt.Println("hello")
	// hello
	// world
}
