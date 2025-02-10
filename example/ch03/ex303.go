package main

import "fmt"

func main() {
	fmt.Println("=== 3.3 文字列、rune、バイト ===")
	{
		var s string = "Hello there"
		var b byte = s[6]
		fmt.Printf("%s", string(b)) // t
	}
	{
		// スライス式も使える
		var s string = "Hello there"
		var s2 string = s[4:7] // o t
		var s3 string = s[:5]  // Hello
		var s4 string = s[6:]  // there
		fmt.Println(s2)
		fmt.Println(s3)
		fmt.Println(s4)
	}
}
