package main

import (
	"fmt"
	"time"
)

// goルーチンサイズ数
//var gonum = 100000000 ストール
//var gonum = 10000000 // きつい

//var gonum = 1000000 918.296msec
//var gonum = 100000 95.0491msec
var gonum = 1000000

// 送信データサイズ数.
var sendsize = 1

func main() {

	ch := make([]chan int, gonum)
	for i, _ := range ch {
		ch[i] = make(chan int)
	}

	fmt.Printf("生成!!\n")

	start := time.Now()
	go first(ch[0])

	for i := 0; i < gonum-1; i++ {
		go pipe(ch[i+1], ch[i])
		//	fmt.Printf("チャネル生成 %d, %d\n", i+1, i)
	}

	fmt.Printf("送信\n")

	printer(ch[gonum-1])
	fmt.Printf("経過時間 %s", time.Since(start))
}

func first(out chan<- int) {
	for i := 0; i < sendsize; i++ {
		out <- i
	}

	close(out)
}

func printer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}

func pipe(out chan<- int, in <-chan int) {
	for v := range in {
		out <- v
	}
	//fmt.Printf("close\n")
	close(out)
}
