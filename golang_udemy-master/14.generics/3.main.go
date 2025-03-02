package main

import "fmt"

type Number interface {
	// ~(チルダ)を付けると独自型の元となる型が一致していればOKとなる
	// 今回の例だとMyInt型はintなのでMax関数でMyInt型の引数を渡せるようになる
	~int | ~int32 | ~int64 | ~float32 | ~float64
}

func Max[T Number](x, y T) T {
	if x > y {
		return x
	}
	return y
}

// interfaceで定義せずにインラインで書くこともできる
func Max2[T ~int | ~int32 | ~int64 | ~float32 | ~float64](x, y T) T {
	if x > y {
		return x
	}
	return y
}

type MyInt int

func main() {
	// Max()
	fmt.Println(Max[int](1, 2))
	fmt.Println(Max[float64](1.1, 1.3))

	var x, y MyInt = 1, 2
	fmt.Println(Max[MyInt](x, y))

	// Max2()
	fmt.Println(Max2[int](1, 2))
	fmt.Println(Max2[float64](1.1, 1.3))

	var x2, y2 MyInt = 1, 2
	fmt.Println(Max2[MyInt](x2, y2))
}
