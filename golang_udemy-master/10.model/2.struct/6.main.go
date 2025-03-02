package main

import "fmt"

type User6 struct {
	Name string
	Age  int
}

func main() {
	m := map[int]User6{
		1: {Name: "user1", Age: 10},
		2: {Name: "user2", Age: 20},
	}
	fmt.Println(m)
	for _, v := range m {
		fmt.Println(v)
	}

	m2 := map[User6]string{
		{Name: "user1", Age: 10}: "Tokyo",
		{Name: "user2", Age: 20}: "LA",
	}
	fmt.Println(m2)

	m3 := make(map[int]User6)
	fmt.Println(m3)
	m3[1] = User6{Name: "user3"}
	fmt.Println(m3)
}
