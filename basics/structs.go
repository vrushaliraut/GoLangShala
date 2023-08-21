package main

import "fmt"

type ninja struct {
	name  string
	level int
}

func structsDemo() {
	fmt.Println("Declare struct ::", ninja{name: "Riyanshi", level: 1})
	fmt.Println("Declare struct ::", ninja{name: "Abhi", level: 2})

	//default values to struct get assigned
	fmt.Println("Declare struct : :name ", ninja{
		name: "john",
	})
}
