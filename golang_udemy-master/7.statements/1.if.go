package main

import "fmt"

func main() {
	a := 2
	if a == 2 {
		fmt.Println("two")
	} else if a == 1 {
		fmt.Println("one")
	} else {
		fmt.Println("I do not know")
	}

	if b := 100; b == 100 {
		fmt.Println("one houndred")
	}

	// 変数のスコープ
	x := 0
	if x := 2; true {
		fmt.Println(x)
	}
	fmt.Println(x)
}
