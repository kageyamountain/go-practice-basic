package main

import "fmt"

func main() {
	i := 0
	for {
		i++
		if i == 3 {
			break
		}
		fmt.Println("Loop")
	}

	point := 0
	for point < 10 {
		fmt.Println(point)
		point++
	}

	for i := 0; i < 10; i++ {
		if i == 3 {
			continue
		}

		if i == 6 {
			break
		}

		fmt.Println(i)
	}

	arr := [3]int{1, 2, 3}
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	// 所謂foreach
	for i, v := range arr {
		fmt.Println(i, v)
	}

	slice := []string{"Python", "PHP", "GO"}
	for i, v := range slice {
		fmt.Println(i, v)
	}

	m := map[string]int{"apple": 100, "banana": 200}
	for k, v := range m {
		fmt.Println(k, v)
	}

}
