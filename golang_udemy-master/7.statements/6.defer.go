package main

import (
	"fmt"
	"os"
)

func TestDefer() {
	defer fmt.Println("END")
	fmt.Println("START")
}

// deferが複数回定義されてる場合は後から登録されたものが順番に実行される
func RunDefer() {
	defer fmt.Println("A")
	defer fmt.Println("B")
	defer fmt.Println("C")
}

func main() {
	TestDefer()

	defer func() {
		fmt.Println("1")
		fmt.Println("2")
		fmt.Println("3")
	}()

	RunDefer()

	file, err := os.Create("test.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	file.Write([]byte("Hello"))
}
