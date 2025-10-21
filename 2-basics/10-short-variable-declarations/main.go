package main

import "fmt"

func main() {
	var i, j int = 1, 2
	k := 3 // :=は関数内でのみ利用可能, :=は自動で型推論してくれる・varが必要ない
	c, python, java := true, false, "no!" // 複数定義することも可能

	fmt.Println(i, j, k, c, python, java) // 1 2 3 true false no!
}