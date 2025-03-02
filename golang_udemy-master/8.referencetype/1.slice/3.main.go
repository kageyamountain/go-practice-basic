package main

import "fmt"

func main() {
	// 参照型なので元のsl[0]も変更される
	sl := []int{100, 200}
	sl2 := sl
	sl2[0] = 1000
	fmt.Println(sl)

	// 基本型なので参照型とは異なり元のiは変更されない
	i := 10
	i2 := i
	i2 = 100
	fmt.Println(i, i2)

	// スライスのコピー
	sl3 := []int{1, 2, 3, 4, 5}
	sl4 := make([]int, 5, 10)
	fmt.Println(sl4)
	n := copy(sl4, sl3) // n = コピーした数を表す
	fmt.Println(n, sl4)
}
