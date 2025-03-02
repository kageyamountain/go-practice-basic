package main

import "fmt"

type User4 struct {
	Name string
	Age  int
}

// コンストラクタ
// 言語仕様としてのコンストラクタはないがこういう形でポインタ型を返すパターンで実装することが多い
func NewUser(name string, age int) *User4 {
	return &User4{Name: name, Age: age}
}

func main() {
	user1 := NewUser("user1", 10)
	fmt.Println(user1)  // 参照を出力
	fmt.Println(*user1) // 実態を出力
}
