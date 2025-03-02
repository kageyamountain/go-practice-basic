package main

import "fmt"

func main() {
	f := func(x, y int) int {
		return x + y
	}
	fmt.Println(f(1, 2))

	i2 := func(x, y int) int {
		return x + y
	}(1, 2)
	fmt.Println(i2)
}
