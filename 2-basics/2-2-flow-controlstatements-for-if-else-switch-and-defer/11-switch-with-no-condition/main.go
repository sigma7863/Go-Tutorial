package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	switch { // switchと{}の間の条件式は何と比較するかを表す、条件式がないとswitch true {}になり、caseに書かれた式が当てはまる
	case t.Hour() < 12:
		fmt.Println("Good morning!") // 実行日時: 2025/10/27(月) 5:09
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}
