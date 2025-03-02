package main

import "fmt"

/*
func 関数名(引数　引数の型) 戻り値型 {
	関数の中身
	return 返す値
}
*/

func Plus(x, y int) int {
	return x + y
}

func Div(x, y int) (int, int) {
	q := x / y
	r := x % y
	return q, r
}

func Double(price int) (result int) {
	result = price * 2
	return
}

func NoReturn() {
	fmt.Println("No Return")
}

func main() {
	i := Plus(1, 2)
	fmt.Println(i)

	i2, i3 := Div(9, 3)
	fmt.Println(i2, i3)

	i4 := Double(1000)
	fmt.Println(i4)

	NoReturn()
}
