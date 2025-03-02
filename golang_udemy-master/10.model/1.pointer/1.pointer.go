package main

import "fmt"

func Double(i int) {
	i = i * 2
}

func Double2(i *int) {
	*i = *i * 2
}

func Double3(s []int) {
	for i, v := range s {
		s[i] = v * 2
	}
}

func main() {
	var n int = 100
	fmt.Println(n)
	fmt.Println(&n)

	Double(n)
	fmt.Println(n)

	// &はポインタを表す、*は実態を表す
	var p *int = &n
	fmt.Println(p)
	fmt.Println(*p)

	*p = 300
	fmt.Println(n, *p)
	n = 200
	fmt.Println(n, *p)

	Double2(&n)
	fmt.Println(n)

	Double2(p)
	fmt.Println(*p)

	// スライスは参照型なのでポインタを使わずとも更新される
	var sl []int = []int{1, 2, 3}
	Double3(sl)
	fmt.Println(sl)
}
