package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	// Print the PID of the current process
	pid := os.Getpid()
	fmt.Printf("PID: %d\n", pid)

	for i := 0; i < 60; i++ {
		fmt.Printf("Running... %d seconds elapsed\n", i)
		time.Sleep(1 * time.Second)
	}
	fmt.Println("Finished running for 1 minute.")
}
