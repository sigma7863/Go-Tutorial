package main

import "fmt"

func main() {
	sum := 1
	// 初項1, 公比2, 末項1024の等比数列
	for ; sum < 1000; { // 初期化と後処理ステートメントは書かなくてもいい
		sum += sum
	}
	fmt.Println(sum) // 1024
}
