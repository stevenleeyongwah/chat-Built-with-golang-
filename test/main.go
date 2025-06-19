package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	buffer      = []int{}
	bufferLimit = 5
	lock        = sync.Mutex{}
	cond        = sync.NewCond(&lock)
)

func producer(id int) {
	for i := 0; i < 10; i++ {
		cond.L.Lock()
		for len(buffer) == bufferLimit {
			fmt.Println("Producer", id, "waiting...")
			cond.Wait() // Wait until there's space in buffer
		}

		buffer = append(buffer, i)
		fmt.Printf("Producer %d produced: %d\n", id, i)
		fmt.Println("Buffer:", buffer)
		cond.Signal() // Wake up one waiting consumer
		cond.L.Unlock()
		time.Sleep(time.Millisecond * 100)
	}
}

func consumer(id int) {
	for {
		cond.L.Lock()
		for len(buffer) == 0 {
			fmt.Println("Consumer", id, "waiting...")
			cond.Wait() // Wait until buffer has items
		}

		item := buffer[0]
		buffer = buffer[1:]
		fmt.Printf("Consumer %d consumed: %d\n", id, item)
		cond.Signal() // Wake up one waiting producer
		cond.L.Unlock()
		time.Sleep(time.Millisecond * 150)
	}
}

func main() {
	go producer(1)
	go producer(2)
	// go consumer(1)

	// Let goroutines run for a while
	time.Sleep(500 * time.Second)
}
