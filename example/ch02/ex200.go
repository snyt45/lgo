package main

import "fmt"

func main() {
	fmt.Println("=== 2.1 基本型 ===")
	fmt.Println("=== 2.1.1 ゼロ値 ===")
	{
		var a int
		var b float32
		var c string
		var d bool

		fmt.Println("a: ", a)
		fmt.Println("b: ", b)
		fmt.Println("c: ", c)
		fmt.Println("d: ", d)
	}

	fmt.Println("=== 2.1.2 リテラル ===")
	fmt.Println("=== 2.1.2.1 整数リテラル ===")
	{
		fmt.Println(1_000_000)
	}

	fmt.Println("=== 2.1.2.2 浮動小数点数リテラル ===")
	{
		fmt.Println(3.14)
	}

	fmt.Println("=== 2.1.2.3 rune リテラル ===")
	{
		fmt.Println('a')
		fmt.Println('\141')
		fmt.Println('\n')
		fmt.Println('\t')
	}

	fmt.Println("=== 2.1.2.4 文字列 リテラル ===")
	{
		x := `バッククォートを使った"ロー文字列リテラル"を使うことで
改行や二重引用符（ダブルクォート）を文字列中に書ける`
		fmt.Println(x)
	}

	fmt.Println("=== 2.1.2.5 リテラルと型 ===")
	{
		var x float32 = 2 * 3.14 // 型が違うがエラーにならない
		fmt.Println(x)
		var y float32 = 2 // 型が違うがエラーにならない
		fmt.Println(y)

		// var a int = "abc" // コンパイル時のエラー
		// var b byte = 1000 // コンパイル時のエラー
	}

	fmt.Println("=== 2.1.3 論理型 ===")
	{
		var flag bool
		fmt.Println(flag) // ゼロ値はfalse
	}

	fmt.Println("=== 2.1.4 数値型 ===")
	fmt.Println("=== 2.1.4.1 整数型 ===")
	{
		var a int
		var b uint
		fmt.Println(a, b) // 整数型のゼロ値は0
	}

	fmt.Println("=== 2.1.4.4 整数関連の演算子 ===")
	{
		var x int = 10
		x *= 2
		fmt.Println(x)
	}

	fmt.Println("=== 2.1.4.5 浮動小数点数型 ===")
	{
		var x float64
		x = 1.2345
		// x % 3 // floatに%は使えない

		fmt.Println("x/0=", x/0) // +Inf
		fmt.Println("-x/0=", -x/0) // -Inf
	}

	fmt.Println("=== 2.1.5 文字列型とrune型 ===")
	{
		var a, b string
		fmt.Println(a,b) // ゼロ値は空文字列
		a = "あいう"
		b = "えお"
		fmt.Println(a+b)
	}

	fmt.Println("=== 2.2 変数の宣言 ===")
	{
		var a int = 10
		var b = 10
		var c int
		var d, e int
		var f, g = 10, "hello"
		fmt.Println(a, b, c, d, e, f, g)
	}
	{
		var (
			a int = 10
			b = 10
			c int
			d, e int
			f, g = 10, "hello"
		)
		fmt.Println(a, b, c, d, e, f, g)
	}
	{
		x := 10
		fmt.Println(x)
		y, z := 30, "hello"
		fmt.Println(y, z)
	}
}
