package main

import (
	"sync"
	"fmt"
)

func producer(queue chan string, waitGroup *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		fmt.Println("producing the item: ", (i+1))
		queue <- fmt.Sprintf("item: %d", (i+1))
	}
	close(queue)
	waitGroup.Done()
}

func consumer(queue chan string, waitGroup *sync.WaitGroup) {
	for val := range queue {
		fmt.Println("consuming the ", val)
	}
	waitGroup.Done()
}

func main() {
	fmt.Println("Yet Another Producer, Consumer Solution !!! YAPCS")
	queue := make(chan string)
	var waitGroup sync.WaitGroup
	waitGroup.Add(2)
	go producer(queue, &waitGroup)
	go consumer(queue, &waitGroup)
	waitGroup.Wait()
	fmt.Println("-------------------------------------------------")
}