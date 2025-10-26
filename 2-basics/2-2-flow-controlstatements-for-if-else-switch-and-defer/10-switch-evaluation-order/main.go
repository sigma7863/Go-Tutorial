package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday() // 今の曜日を確かめる
	// switchのcaseは上から下へ評価する、caseの条件が一致すれば、そこで停止(自動的にbreak)する
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.") // 実行したのが2025/10/26(日曜日)だったので、defaultが出力された
	}
}
