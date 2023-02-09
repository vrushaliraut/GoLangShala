package main

import (
	"fmt"
	"time"
)

func attack3(target string, attacked chan bool) {
	time.Sleep(time.Second)
	fmt.Println("Throwing ninja stars at ", target)
	attacked <- true
}

func channelsDemo() {
	smokeSignal := make(chan bool)
	evilNinja := "Tommy"
	go attack3(evilNinja, smokeSignal)

	fmt.Println(<-smokeSignal)

	//buffered channel
	moreSmokeSignal := make(chan bool, 1)
	moreSmokeSignal <- true
	fmt.Println(<-moreSmokeSignal)

	//closing channel to prevent deadlock
	moreSmokeSignal <- true
	close(moreSmokeSignal)
	for message := range moreSmokeSignal {
		fmt.Println(message)
	}
}
