package main

import "fmt"

func main() {
	defer func() {
		if x := recover(); x != nil {
			fmt.Fprintln(x)
		}
	}()
	panic("runtime error")
	fmt.Println("Start")
}
