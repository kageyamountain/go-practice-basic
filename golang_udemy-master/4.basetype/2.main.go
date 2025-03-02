package main

import "fmt"

func main() {
	var i int = 100
	fmt.Println(i)

	var i2 int64 = 200
	fmt.Println(i2)
	fmt.Println(i2 + 50)

	// データ型の確認
	fmt.Printf("%T\n", i2)

	// 型変換
	fmt.Printf("%T\n", int32(i2))

	fmt.Println(i + int(i2))
}
