package main

import (
	"fmt"
	"reflect"
)

func main() {
	// 配列の大きさと要素の型を指定する（値は指定していないのでゼロ値になる）
	var x1 [3]int
	fmt.Println(x1)

	// 配列の大きさと要素の型と値を指定する
	var x2 = [3]int{10, 20, 30}
	fmt.Println(x2)

	// 途中に「空き」がある配列の値を指定する
	var x3 = [12]int{1, 5: 4, 6, 10: 100, 15}
	fmt.Println(x3)

	var x4 = [10]int{1, 2, 3: 9, 6: 7}
	fmt.Println(x4)

	// 配列の初期化に配列リテラルを使う場合は、大きさの指定に ... が使える
	var x5 = [...]int{10, 20, 30}
	fmt.Println(x5)

	// 多次元配列のように扱う
	// x は 長さ2の配列で、その要素の型は長さ3の整数配列
	var x6 [2][3]int
	fmt.Println(x6)

	// 配列の長さを調べる
	fmt.Println(len(x6))

	// Goでは配列の大きさも型の一部になる
	var x7 [3]int
	var x8 [4]int
	// 配列の大きさが違うと型も異なる
	fmt.Println(reflect.TypeOf(x7) != reflect.TypeOf(x8)) // true
}
