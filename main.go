package main

import (
	"fmt"
	"time"
)

func producer(queue chan string) {
	for i := 0; i < 10; i++ {
		fmt.Println("producing the item: ", (i+1))
		queue <- fmt.Sprintf("item: %d", (i+1))
	}
}

func consumer(queue chan string) {
	for val := range queue {
		fmt.Println("consuming the ", val)
	}
}

func main() {
	fmt.Println("Yet Another Producer, Consumer Solution !!! YAPCS")
	queue := make(chan string)
	go producer(queue)
	go consumer(queue)
	time.Sleep(20 * time.Millisecond)
	fmt.Println("-------------------------------------------------")
}