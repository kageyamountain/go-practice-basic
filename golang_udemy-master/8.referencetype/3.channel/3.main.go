package main

import (
	"fmt"
	"time"
)

// channel
// close

func reciever2(name string, ch chan int) {
	for {
		i, ok := <-ch
		if !ok {
			break
		}
		fmt.Println(name, i)
	}
	fmt.Println(name + "END")
}

func main() {
	ch1 := make(chan int, 2)
	close(ch1)

	// closeしたチャネルへ送信はできない
	// ch1 <- 1

	// ただし、closiしたチャネルからの受信はできる
	fmt.Println(<-ch1)

	// ok = チャネルの状態を表す真偽値
	// チャネルが空で、closeされている場合はfalseになる
	i, ok := <-ch1
	fmt.Println(i, ok)

	ch2 := make(chan int, 2)
	go reciever2("1.goroutine", ch2)
	go reciever2("2.goroutine", ch2)
	go reciever2("3.goroutine", ch2)

	x := 0
	for x < 100 {
		ch2 <- x
		x++
	}
	close(ch2)
	time.Sleep(3 + time.Second)
}
