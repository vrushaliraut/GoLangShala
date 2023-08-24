package main

import (
	"fmt"
	"sync"
	"time"
)

func func_go_routine_thread() {
	var wait_group sync.WaitGroup

	//goroutines
	for i := 1; i <= 3; i++ {
		wait_group.Add(1)
		go func(id int) {
			defer wait_group.Done()
			fmt.Printf("Goroutine %d started.\n", id)
			time.Sleep(1 * time.Second)
			fmt.Printf("Goroutine %d completed.\n", id)
		}(i)
	}
	wait_group.Wait()
}
