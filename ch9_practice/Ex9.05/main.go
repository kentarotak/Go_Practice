package main

import (
	"fmt"
	"time"
)

// 会社のPC:4798371回

var counter int

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	ticker := time.NewTicker(1 * time.Second)

	go func() {
		ch1 <- 0
	}()

	go ch_rcv(ch2, ch1)
	go ch_rcv(ch1, ch2)

	<-ticker.C
	fmt.Printf("通信回数 %d \n", counter)
	close(ch1)
	close(ch2)

}

func ch_rcv(out chan<- int, in <-chan int) {
	for {
		counter = <-in
		out <- (counter + 1)
	}
}
