package main

import "fmt"

func main() {
	sum := 1
	// 初項1, 公比2, 末項1024の等比数列(1、2、4、8、16...)
	// Goはforのみでwhileはない
	for sum < 1000 { // forとsum < 1000の間の;(セミコロン)は省略可能
		sum += sum
	}
	fmt.Println(sum) // 1024
}
