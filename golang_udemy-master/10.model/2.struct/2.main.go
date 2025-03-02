package main

import "fmt"

type User2 struct {
	Name string
	Age  int
}

func (u User2) SayName() {
	fmt.Println(u.Name)
}

// レシーバがポインタ型ではないので更新が呼び出し元に反映されない
func (u User2) SetName(name string) {
	u.Name = name
}

// レシーバがポインタ型なので更新が呼び出し元に反映される
func (u *User2) SetName2(name string) {
	u.Name = name
}

func main() {
	user1 := User2{Name: "user1"}
	user1.SayName()

	user1.SetName("A")
	user1.SayName()

	user1.SetName2("A")
	user1.SayName()

	user2 := &User2{Name: "user2"}
	user2.SetName2("B")
	user2.SayName()
}
