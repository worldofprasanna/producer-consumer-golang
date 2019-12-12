package main

import (
	"time"
	"os"
	"os/signal"
	"sync"
	"fmt"
	"syscall"
)

func producer(queue chan string, waitGroup *sync.WaitGroup, quit <-chan bool) {
	defer close(queue)
	for i := 0; i < 10; i++ {
		select {
		case <-quit:
			waitGroup.Done()
			return
		default:
			fmt.Println("producing the item: ", (i+1))
			queue <- fmt.Sprintf("item: %d", (i+1))
		}
	}
	waitGroup.Done()
}

func consumer(queue chan string, waitGroup *sync.WaitGroup) {
	for val := range queue {
		time.Sleep(time.Second)
		fmt.Println("consuming the ", val)
	}
	waitGroup.Done()
}

func handleSigInt(sigInt chan os.Signal, queue chan string, quit chan<- bool) {
	// Wait for the Signal Interrupt
	_ = <- sigInt
	fmt.Println("Consume the pending items and terminating gracefully")
	quit <- true
}

func main() {
	fmt.Println("Yet Another Producer, Consumer Solution !!! YAPCS")
	queue := make(chan string)
	sigInt, quit := make(chan os.Signal), make(chan bool)
	signal.Notify(sigInt, syscall.SIGINT, syscall.SIGTERM)

	// Handle the SIGINT signal as a separate goroutine
	go handleSigInt(sigInt, queue, quit)

	var waitGroup sync.WaitGroup
	waitGroup.Add(2)
	go producer(queue, &waitGroup, quit)
	go consumer(queue, &waitGroup)

	waitGroup.Wait()
	fmt.Println("-------------------------------------------------")
}