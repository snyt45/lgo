package main

import "fmt"

func main() {
	fmt.Println("=== 2.4 型付きの定数と型のない定数 ===")
	// 型なしの定数を宣言
	const x = 10

	// 数値型に柔軟に代入できる
	var y int = x
	var z float64 = x
	var d byte = x

	fmt.Println(y, z, d)

	// 型付きの定数を宣言
	const typedX int = 10

	// var z float64 = typedX // コンパイルエラーになる
}
