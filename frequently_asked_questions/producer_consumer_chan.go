package main

import (
	"fmt"
	"sync"
)

func func_multiple_goroutines_with_chan() {
	channel := make(chan int)
	
	var wait_group sync.WaitGroup

	wait_group.Add(2)
	go producer(channel, &wait_group)
	go consumer(channel, &wait_group)

	wait_group.Wait()
}

func producer(channel chan int, wait_group *sync.WaitGroup) {
	defer wait_group.Done()
	for i := 1; i <= 5; i++ {
		channel <- i
	}
	close(channel)
}

func consumer(channel <-chan int, wait_group *sync.WaitGroup) {
	defer wait_group.Done()
	for num := range channel {
		fmt.Printf("Received: %d \n", num)
	}
}
