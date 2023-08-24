package main

import "fmt"

// function
func greet(name string) {
	fmt.Printf("Hello, %s ! \n", name)
}

// Method
type Person struct {
	Name string
}

func (p Person) introduce() {
	fmt.Printf("My name is, %s ! \n", p.Name)
}
