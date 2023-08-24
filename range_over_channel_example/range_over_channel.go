package main

import "fmt"

//https://techwasti.com/range-over-channel-in-go-lang

func main() {

	/*
		- create channel
		- send string value to channel
		- use range to receive values from channel
	*/

	fmt.Println("*** range with channel *** \n ")
	ch := make(chan string)
	go func() {
		ch <- "Hello"
		ch <- "world"
		ch <- "!"
		close(ch)
	}()

	for msg := range ch {
		fmt.Println(msg)
	}

	/*
		- create buffered channel
		- send string value to channel
		- use range to receive values from channel
	*/

	channel := make(chan int, 2)
	go func() {
		channel <- 1
		channel <- 2

		close(channel)
	}()

	for msg := range channel {
		fmt.Println(msg)
	}

	// range with Arrays, slice , maps, strings, and channel we just saw above
	range_in_go()
}
