package main

import "fmt"

var i5 int = 500

func outer() {
	var s4 string = "outer"
	fmt.Println(s4)
}

func main() {
	// 明示的な変数定義
	// var 変数名 型 = 値
	var i int = 100
	fmt.Println(i)

	var s string = "Hello Go"
	fmt.Println(s)

	var t, f bool = true, false
	fmt.Println(t, f)

	var (
		i2 int    = 200
		s2 string = "Golang"
	)
	fmt.Println(i2, s2)

	var i3 int
	var s3 string
	fmt.Println(i3, s3) // 初期値int=0, string=""が出力される

	i3 = 300
	s3 = "Go"
	fmt.Println(i3, s3)

	i = 150
	fmt.Println(i)

	// 暗黙的な変数定義
	// 変数名 := 値
	i4 := 400
	fmt.Println(i4)

	i4 = 450
	fmt.Println(i4)

	fmt.Println(i5)

	outer()
}

// import "fmt"

// var i int = 100
// var s string = "Golang"

// var t, f bool = true, false

// var (
// 	ii int    = 1000
// 	ss string = "Go"
// )

// var i2 int

// var sss string

// //i3 := 100

// func main() {
// 	fmt.Println(i2)
// 	i2 = 150
// 	fmt.Println(i2)

// 	fmt.Println(sss)
// 	sss = "Go!!"
// 	fmt.Println(sss)

// 	i2 = 200
// 	fmt.Println(i2)

// 	i3 := 200
// 	fmt.Println(i3)

// 	//i3 := 300

// 	i3 = 300

// 	//i3 = "string"

// }
