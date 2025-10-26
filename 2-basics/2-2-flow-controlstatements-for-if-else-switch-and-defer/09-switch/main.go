package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("Go runs on ") // onの後にスペースを入れないとonの後の部分にスペースが入らない
	// 使っているOSを判別する
	switch os := runtime.GOOS; os {
	// caseの最後にbreakは必要ない(自動でやってくれる)、ほかの言語と違ってswitch文のcaseに書く値は変数でも定数でも何でもいい
	case "darwin":
		fmt.Println("OS X.") // macOSを使っていると出力される
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os) // windowsの場合: Go runs on windows.
	}
}
