package main

import (
	"fmt"
	"time"
)

func attack1(target string) {
	fmt.Println("Throwing ninja stars at ", target)
}

func go_routine_demo() {
	go attack1("Tommy")
	time.Sleep(time.Second)

	go printNumbers() // start a goroutine
	go printLetters() // another goroutine

	time.Sleep(2 * time.Second) //wait for goroutine to finish
}

// Example of using gorouitnes to perform concrrent tasks
func printNumbers() {
	for i := 1; i <= 5; i++ {
		fmt.Printf("%d", i)
		time.Sleep(100 * time.Millisecond)
	}
}

func printLetters() {
	for i := 'a'; i <= 'e'; i++ {
		fmt.Printf("%c", i)
		time.Sleep(200 * time.Millisecond)
	}
}
