package main

import (
	"fmt"
	"strconv"
)

func main() {
	var i int = 1
	fl64 := float64(i)
	fmt.Println(fl64)
	fmt.Printf("fl64 = %T\n", fl64)

	i2 := int(fl64)
	fmt.Printf("i2 = %T\n", i2)

	fl := 10.5
	i3 := int(fl)
	fmt.Printf("i3 = %T\n", i3)
	fmt.Println(i3)

	var s string = "100"
	fmt.Printf("s = %T\n", s)

	// x, err := strconv.Atoi(s)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	x, _ := strconv.Atoi(s)
	fmt.Println(x)

	var x2 int = 200
	s2 := strconv.Itoa(x2)
	fmt.Println(s2)
	fmt.Printf("s2 = %T\n", s2)

	var h string = "Hello World"
	b := []byte(h)
	fmt.Println(b)

	h2 := string(b)
	fmt.Println(h2)
}
