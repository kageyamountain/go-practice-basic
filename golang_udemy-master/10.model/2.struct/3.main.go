package main

import "fmt"

type T struct {
	User User3
}

type User3 struct {
	Name string
	Age  int
}

func (u *User3) SetName(name string) {
	u.Name = name
}

func main() {
	t := T{
		User: User3{
			Name: "user1",
			Age:  10,
		},
	}
	fmt.Println(t)
	fmt.Println(t.User)
	fmt.Println(t.User.Name)
	fmt.Println(t.User.Age)

	t.User.SetName("A")
	fmt.Println(t)
}
