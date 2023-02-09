package main

import (
	"fmt"
	"time"
)

func attack1(target string) {
	fmt.Println("Throwing ninja stars at ", target)
}
func goRoutineDemo() {
	go attack1("Tommy")
	time.Sleep(time.Second)
}
