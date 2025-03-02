package main

import "fmt"

func anything(a interface{}) {
	switch v := a.(type) {
	case string:
		fmt.Println(v + "!?")
	case int:
		fmt.Println(v + 10000)
	}
}

func main() {
	anything(1)
	anything("AAA")

	var x interface{} = 3
	i := x.(int) // 型アサーション
	fmt.Println(i + 2)

	f, isFloat64 := x.(float64)
	fmt.Println(f, isFloat64)

	// if文でやるのは冗長な記述
	if x == nil {
		fmt.Println("None")
	} else if i, isInt := x.(int); isInt {
		fmt.Println(i, "x is Int")
	} else if s, isString := x.(string); isString {
		fmt.Println(s, "x is string")
	} else {
		fmt.Println("others")
	}

	// 型switch
	switch x.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("others")
	}

	// 型switchで値を使う方法
	switch v := x.(type) {
	case int:
		fmt.Println(v, "int")
	case string:
		fmt.Println(v, "string")
	default:
		fmt.Println(v, "others")
	}
}
