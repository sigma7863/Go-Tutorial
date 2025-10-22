package main

import "fmt"

func main() {
	sum := 0
	// 1~10までの和を求める(初項1, 公差1, 末項10)
	for i := 0; i < 10; i++ { // 初期化ステートメント, 条件式, 後処理ステートメントの順に書く
		sum += i
	}
	// for (〇〇) {}ではなくfor 〇〇 {}と書く, ()は必要ない
	fmt.Println(sum) // 45
}
