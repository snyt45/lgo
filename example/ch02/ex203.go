package main

import "fmt"

const x int64 = 10

const (
	idKey = "id"
	nameKey = "name"
)

const z = 20 * 10

func main() {
	const y = "hello"

	fmt.Println(x)
	fmt.Println(y)

	// x = x + 1 // 定数はイミュータブルなので変更できない
	// y = "bye" // 定数はイミュータブルなので変更できない
	// 
	// fmt.Println(x)
	// fmt.Println(y)
}
