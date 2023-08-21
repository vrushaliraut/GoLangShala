package main

import "fmt"

func ranges() {
	evilNinjas := []string{"Riyu", "Bhagyashri", "Vrushali"}
	for index, evilNinja := range evilNinjas {
		fmt.Println("Attacking target", index, evilNinja)
	}

	evilNinjasWithLevel := map[string]int{"Abhi": 2}

	for evilNinja, level := range evilNinjasWithLevel {
		fmt.Println("%s -> %d \n", evilNinja, level)
	}
}
