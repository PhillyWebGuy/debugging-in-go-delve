package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println(t)

	f, err := os.Open("./readme.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s", err)
		os.Exit(1)
	}
	defer f.Close()

	data, err := io.ReadAll(f)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s", err)
		os.Exit(1)
	}
	fmt.Println(string(data))
}
