package main

import (
	"fmt"
	"time"
)

func func_select_switch() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- 42
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- 23
	}()

	select {
	case val := <-ch1:
		fmt.Println("Received from ch1:", val)
	case val := <-ch2:
		fmt.Println("Received from ch2:", val)
	}

	num := 2
	switch num {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Switch statement :: Two")
	default:
		fmt.Println("Other")
	}

}
