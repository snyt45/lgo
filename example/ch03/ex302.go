package main

import "fmt"

func main() {
	fmt.Println("=== 3.2 スライス ===")
	{
		// 宣言時に大きさを指定しない
		// [n]あるいは[...]と書くと配列になる
		var x1 = []int{10, 20, 30}
		fmt.Println(x1)

		// インデックスを指定して要素の一部だけを指定できる
		var x2 = []int{1, 5: 4, 10: 100, 15}
		fmt.Println(x2)

		// 多次元のスライス
		var x3 [][]int
		fmt.Println(x3)

		// スライスのゼロ値。nilが初期値になる
		var x4 []int
		fmt.Println(x4)

		// スライス同士の比較はできない
		// コンパイルエラーになる
		// fmt.Println(x3 == x4)

		// nilと比較は可能
		fmt.Println(x4 == nil)
	}

	fmt.Println("=== 3.2.1 len ===")
	{
		// 配列の長さを調べる
		var x1 = [2]int{10, 20}
		fmt.Println(len(x1))

		// 配列リテラルを未設定
		var x2 [2]int
		fmt.Println(len(x2))

		// スライスの長さを調べる
		var x3 = []int{5, 10}
		fmt.Println(len(x3))

		// nilスライスの場合
		var x4 []int
		fmt.Println(len(x4))
	}

	fmt.Println("=== 3.2.2 append ===")
	{
		var x []int
		x = append(x, 10)
		fmt.Println(x)

		// 複数の値の追加
		x = append(x, 5, 6, 7)
		fmt.Println(x)

		// 演算子 ... スライスの個々の値を展開できる
		y := []int{20, 30, 40}
		x = append(x, y...)
		fmt.Println(x)
	}

	fmt.Println("=== 3.2.3 スライスのキャパシティ ===")
	{
		var x []int
		fmt.Println(x, len(x), cap(x))
		x = append(x, 10)
		fmt.Println(x, len(x), cap(x))
		x = append(x, 20)
		fmt.Println(x, len(x), cap(x))
		x = append(x, 30)
		fmt.Println(x, len(x), cap(x))
		x = append(x, 40)
		fmt.Println(x, len(x), cap(x))
		x = append(x, 50)
		fmt.Println(x, len(x), cap(x))
	}

	fmt.Println("=== 3.2.4 make ===")
	{
		// 長さとキャパシティを一括指定
		// この場合、長さ5でキャパシティ5でゼロ値が0のスライスが生成される
		x1 := make([]int, 5)
		fmt.Println(x1)

		// appendすると、6個目の要素に追加される
		x1 = append(x1, 5)
		fmt.Println(x1)

		// 長さとキャパシティをそれぞれ指定
		// 長さ5, キャパシティ10のスライス
		x2 := make([]int, 5, 10)
		x2 = append(x2, 1, 2, 3)
		fmt.Println(x2)

		// 長さ0, キャパシティ10のスライス
		x3 := make([]int, 0, 10)
		x3 = append(x3, 1, 2, 3)
		fmt.Println(x3)
	}

	fmt.Println("=== 3.2.5 スライスの生成方法の選択 ===")
	{
		// nilスライスを宣言
		var data1 []int
		fmt.Println(len(data1), cap(data1))
		fmt.Println(data1 == nil)

		// 空のスライスリテラルを宣言
		data2 := []int{}
		fmt.Println(len(data2), cap(data2))
		fmt.Println(data2 == nil)

		// デフォルト値を宣言
		data3 := []int{2, 4, 6, 8}
		fmt.Println(len(data3), cap(data3))
		fmt.Println(data3 == nil)
	}

	fmt.Println("=== 3.2.6 スライスのスライス ===")
	{
		x := []int{1, 2, 3, 4}
		y := x[:2]
		z := x[1:]
		d := x[1:3]
		e := x[:]
		fmt.Println("x:", x)
		fmt.Println("y:", y)
		fmt.Println("z:", z)
		fmt.Println("d:", d)
		fmt.Println("e:", e)
	}
}
