package main

import "fmt"

func loopsDemo() {
	isSkilled := true
	for isSkilled {
		fmt.Println("Ready")
		isSkilled = false
	}

	for level := 7; level < 9; level++ {
		fmt.Println(level)
		fmt.Println("Leveling up..!")
	}

	for {
		fmt.Println("I'm a java ninja")
		break
	}

}
