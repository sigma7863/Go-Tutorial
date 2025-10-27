package main

import "fmt"

func main() {
	fmt.Println("counting") // この部分はforとdoneの間に配置してもcounting, done, 9, 8...になる

	for i := 0; i < 10; i++ {
		defer fmt.Println(i) // 0, 1, 2...ではなく9, 8, 7...の順で実行される(LIFO(Last-In-First-Out))
	}

	fmt.Println("done")
	// counting
	// done
	// 9
	// 8
	// 7
	// 6
	// 5
	// 4
	// 3
	// 2
	// 1
	// 0
}
