package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Second * 1)

	// fmt.Println("Ticker starting")
	fmt.Fprintf(os.Stdout, "Ticker starting \n")

	go func() {
		for t := range ticker.C {
			fmt.Fprintf(os.Stdout, "Hello world from %v \n", t)
			fmt.Fprintf(os.Stderr, "ERROR at %v \n", t)
		}
	}()

	time.Sleep(time.Hour * 1)
	ticker.Stop()

	// fmt.Println("Ticker stopped")
	fmt.Fprintf(os.Stdout, "Ticker stopping \n")
}