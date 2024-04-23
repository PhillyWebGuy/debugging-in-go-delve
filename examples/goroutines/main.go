package main

import (
	"runtime"
	"sync"
	"time"
)

func doStuff(wg *sync.WaitGroup, id int) {
	time.Sleep(1 * time.Second)
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	for i := range 10 {
		wg.Add(1)
		go doStuff(&wg, i)
	}
	runtime.Breakpoint()
	wg.Wait()
}
