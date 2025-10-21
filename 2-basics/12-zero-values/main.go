package main

import "fmt"

func main() {
	var i int // 初期値として0が入る
	var f float64 // floatも同様に0が入る
	var b bool // 初期値としてfalseが入る
	var s string // 初期値として""が入る
	fmt.Printf("%v %v %v %q\n", i, f, b, s) // 0 0 false ""
}
