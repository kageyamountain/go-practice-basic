package main

import "fmt"

// ジェネリクス無しだとこう実装する必要がある
// 引数の型が異なる度に関数を実装する必要があって冗長なコードになる
//func PrintSliceInts(i []int) {
//	for _, v := range i {
//		fmt.Println(v)
//	}
//}
//
//func PrintSliceStrings(s []string) {
//	for _, v := range s {
//		fmt.Println(v)
//	}
//}
//
//func main() {
//	PrintSliceInts([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
//	PrintSliceStrings([]string{"a", "b", "c"})
//}

// ↓ジェネリクスを使った実装↓

func PrintSlice[T any](s []T) {
	for _, v := range s {
		fmt.Println(v)
	}
}

func main() {
	PrintSlice[int]([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	PrintSlice[string]([]string{"a", "b", "c"})
}
