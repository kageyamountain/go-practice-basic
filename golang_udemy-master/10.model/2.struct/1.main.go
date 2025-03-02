package main

import "fmt"

// 構造体
type User struct {
	// フィールド
	Name string
	Age  int
}

func UpdateUser(user User) {
	user.Name = "A"
	user.Age = 1000
}

func UpdateUser2(user *User) {
	user.Name = "A"
	user.Age = 1000
}

func main() {
	var user1 User
	fmt.Println(user1)
	user1.Name = "user1"
	user1.Age = 10
	fmt.Println(user1)

	user2 := User{}
	fmt.Println(user2)
	user2.Name = "user2"
	fmt.Println(user2)

	user3 := User{Name: "user3", Age: 30}
	fmt.Println(user3)

	user4 := User{"user4", 40}
	fmt.Println(user4)

	// これはエラー
	// user5 := User{30, "user5"}

	user6 := User{Name: "user6"}
	fmt.Println(user6)

	// newは構造体のポインタ型を返す
	user7 := new(User)
	fmt.Println(user7)

	// アドレス演算子&を付けるとnewと同じくポインタ型を返す。こっちの方がよく使われる。
	user8 := &User{}
	fmt.Println(user8)

	// 構造体は基本型の扱いなのでポインタ型で渡さないと関数内での変更が反映されない
	UpdateUser(user1)
	UpdateUser2(user8)
	fmt.Println(user1)
	fmt.Println(user8)
}
