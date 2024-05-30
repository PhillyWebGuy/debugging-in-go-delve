package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Println("pid:", os.Getpid())
	chan1 := make(chan time.Time, 3)
	for i := 1; i <= 3; i++ {
		chan1 <- time.Now()
	}

	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			chan1 <- t
		}
	}()

	chan2 := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		chan2 <- i
	}

	for r := range chan2 {
		fmt.Println("received request", r)
	}
}
