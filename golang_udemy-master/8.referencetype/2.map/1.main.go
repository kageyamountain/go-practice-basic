package main

import "fmt"

func main() {
	var m = map[string]int{"A": 100, "B": 200}
	fmt.Println(m)

	m2 := map[string]int{"A": 100, "B": 200}
	fmt.Println(m2)

	m3 := map[int]string{
		1: "A",
		2: "B",
	}
	fmt.Println(m3)

	m4 := make(map[int]string)
	fmt.Println(m4)
	m4[1] = "JAPAN"
	m4[2] = "USA"
	fmt.Println(m4)

	fmt.Println(m["A"])
	fmt.Println(m4[2])

	// 登録されていないキーの場合は初期値が取り出される
	fmt.Println(m4[3])

	// ok = 取り出しに成功したかどうかのtrue/false
	s, ok := m4[1]
	if !ok {
		fmt.Println("Error")
	}
	fmt.Println(s, ok)

	s2, ok2 := m4[3]
	if !ok2 {
		fmt.Println("Error")
	}
	fmt.Println(s2, ok2)

	m4[2] = "US"
	fmt.Println(m4)

	m4[3] = "CHINA"
	fmt.Println(m4)

	// 要素の削除
	delete(m4, 3)
	fmt.Println(m4)

	// 要素数の取得
	fmt.Println(len(m4))

	// map + for
	mf := map[string]int{
		"Apple":  100,
		"Banana": 200,
	}

	for k, v := range mf {
		fmt.Println(k, v)
	}
}
