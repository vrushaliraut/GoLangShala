package main

import "fmt"

func switchDemo() {
	switch_with_string()
	powerlevel := 9001
	conditional_switch(powerlevel)
}

func conditional_switch(powerlevel int) {
	switch {
	case powerlevel > 9000:
		fmt.Println("powerlevel is 9000")
	default:
		fmt.Println("Baby Ninja")
	}
}

func switch_with_string() {
	weapon := "Ninja star"
	switch weapon {
	case "Ninja star":
		fmt.Println("I am a Ninja star")
	case "Ninja sword":
		fmt.Println("I am a Ninja sword")
	}
}
