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

	fmt.Println("=== 3.2.6.1 スライスの記憶領域の共有 ===")
	{
		x := []int{1, 2, 3, 4}
		y := x[:2]
		z := x[1:]

		fmt.Println("x:", x)
		fmt.Println("y:", y)
		fmt.Println("z:", z)

		x[1] = 20
		y[0] = 10
		z[1] = 30

		fmt.Println("x:", x)
		fmt.Println("y:", y)
		fmt.Println("z:", z)
	}
	{
		x := []int{1, 2, 3, 4}
		y := x[:2]

		fmt.Println("x:", x)
		fmt.Println("y:", y)

		// xから別のスライスyを作ると、長さは2でキャパシティは4になる
		fmt.Println(len(x), cap(x), len(y), cap(y)) // 4 4 2 4
		// yの終わりにappendすると、yの終わりとxの3番目の位置に追加される
		y = append(y, 30)

		fmt.Println("x:", x) // x: [1 2 30 4]
		fmt.Println("y:", y) // y: [1 2 30]
	}
	{
		x := make([]int, 0, 5)
		x = append(x, 1, 2, 3, 4)           // x: [1 2 3 4]
		y := x[:2:2]                        // y: [1 2]
		z := x[2:4:4]                       // z: [3 4]
		fmt.Println(cap(x), cap(y), cap(z)) // 5 2 2
		y = append(y, 30, 40, 50)
		x = append(x, 60)
		z = append(z, 70)
		fmt.Println("x:", x) // x: [1 2 3 4 60]
		fmt.Println("y:", y) // y: [1 2 30 40 50]
		fmt.Println("z:", z) // z: [3 4 70]
	}
	fmt.Println("=== 3.2.7 配列からスライスへの変換 ===")
	{
		x := [...]int{5, 6, 7, 8}
		y := x[:2]
		z := x[2:]
		d := x[:]
		x[0] = 10
		fmt.Println("x: ", x) // x:  [10 6 7 8]
		fmt.Println("y: ", y) // y:  [10 6]
		fmt.Println("z: ", z) // z:  [7 8]
		fmt.Println("d: ", d) // d:  [10 6 7 8]
	}
	fmt.Println("=== 3.2.8 メモリを共有しないスライスのコピー ===")
	{
		x := []int{1, 2, 3, 4}
		y := make([]int, 4)
		fmt.Println(y, len(y), cap(y)) // [0 0 0 0] 4 4
		num := copy(y, x)              // xからyにコピー
		fmt.Println(y, num)            // [1 2 3 4] 4
	}
	{
		// スライスの部分のコピー
		// 長さ2のスライスにするとそれ以上はコピーされない
		x := []int{1, 2, 3, 4}
		y := make([]int, 2) // 長さ2のスライスを作る
		num := copy(y, x)   // 先頭から2要素だけコピー
		fmt.Println(y, num) // [1 2] 2
	}
	{
		// スライスの途中をコピー
		// copyの戻り値を変数に代入しなくてもよい
		x := []int{1, 2, 3, 4}
		y := make([]int, 2) // 長さ2のスライスを作る
		copy(y, x[1:])      // x[1]から2コピー
		fmt.Println(y)      // [1 2] 2
	}
	{
		// オーバーラップするスライスのコピー
		x := []int{1, 2, 3, 4}
		num := copy(x[:3], x[1:]) // [1 2 3] を [2 3 4] で上書き
		fmt.Println(x, num)       // [2 3 4 4] 3
	}
	{
		// ターゲットに配列も指定できる
		x := []int{1, 2, 3, 4}
		d := [4]int{5, 6, 7, 8} // d は配列
		y := make([]int, 2)
		copy(y, d[:])  // 配列をcopy
		fmt.Println(y) // [5 6]
		copy(d[:], x)
		fmt.Println(d) // [1 2 3 4]
	}
}
