package main

import "fmt"

func main() {
	sl := []string{"A", "B", "C"}
	fmt.Println(sl)

	// range（foreach的なやつ）
	for i, v := range sl {
		fmt.Println(i, v)
	}

	// 古典的for
	for i := 0; i < len(sl); i++ {
		fmt.Println(i, sl[i])
	}
}
