package main

import "fmt"

type MyInt int

func (mi MyInt) Print() {
	fmt.Println(mi)
}

// 独自型
func main() {
	var mi MyInt
	fmt.Println(mi)
	fmt.Printf("%T\n", mi)

	// これは型が異なるのでエラーになる
	// i := 100
	// fmt.Println(mi + i)

	mi = 1
	mi.Print()
}
