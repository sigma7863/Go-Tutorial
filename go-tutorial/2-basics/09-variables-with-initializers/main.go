package main

import "fmt"

var i, j int = 1, 2 // intで型を明示

func main() {
		var c, python, java = true, false, "no!" // 初期化子(= true, false...の部分, initializer)を定義すると自動で型を推論してくれる
		fmt.Println(i, j, c, python, java) // 1 2 true false no!
}