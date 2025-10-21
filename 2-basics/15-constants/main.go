package main

import "fmt"

const Pi = 3.14

func main() {
	// constは文字(character)、文字列(string)、boolean、数値(numeric)しか使えない
	// :=を使って宣言できない
	const World = "世界"
	fmt.Println("Hello", World) // Hello 世界
	fmt.Println("Happy", Pi, "Day") // Happy 3.14 Day

	const Truth = true
	fmt.Println("Go rules?", Truth) // Go rules? true
}
