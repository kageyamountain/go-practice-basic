package main

import "fmt"

type User5 struct {
	Name string
	Age  int
}

// 構造体スライス
type Users []*User5

func main() {
	user1 := User5{Name: "user1", Age: 10}
	user2 := User5{Name: "user2", Age: 20}
	user3 := User5{Name: "user3", Age: 30}
	user4 := User5{Name: "user4", Age: 40}

	users := Users{}
	users = append(users, &user1)
	users = append(users, &user2)
	users = append(users, &user3, &user4)
	for _, u := range users {
		fmt.Println(*u)
	}

	users2 := make([]*User5, 0)
	users2 = append(users2, &user1, &user2, &user3, &user4)
	for _, u := range users2 {
		fmt.Println(*u)
	}
}
