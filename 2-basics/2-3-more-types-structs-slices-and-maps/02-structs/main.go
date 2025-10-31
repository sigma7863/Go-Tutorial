package main

import "fmt"

type Vertex struct { // type [構造体の名前] struct
	Y int
	X int
}

func main() {
	fmt.Println(Vertex{1, 2}) // {1 2}
}
