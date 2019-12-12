package main

import (
	"fmt"
	"time"
)

func producer() {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Millisecond)
		fmt.Println("producing the item ", i)
	}
}

func consumer() {
}

func main() {
	fmt.Println("Yet Another Producer, Consumer Solution !!! YAPCS")
	// Invoke producer in a separate go routine
	go producer()
	// Invoke consumer in a separate go routine
	go consumer()
	fmt.Println("-------------------------------------------------")
}