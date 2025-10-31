package main

import "fmt"

func main() {
	i, j := 42, 2701


	p := &i 		// point to i...pにiのポインタを割り当て
	fmt.Println(*p) // read i through the pointer...ポインタを通してiを読み出す、42
	*p = 21			// set i through the pointer...ポインタpを通してiへ値を代入する
	fmt.Println(i)  // see the new value of i...iの新しい値を呼び出す...21

	p = &j			// point to j...pにjのポインタを割り当て
	*p = *p / 37.   // divide j through the pointer...ポインタを通してjを割り算する
	fmt.Println(j)	// see the new value of j...jの新しい値を呼び出す...73
}
