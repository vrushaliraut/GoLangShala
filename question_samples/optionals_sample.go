package main

import "fmt"

func explore_different_optionals() {

	fmt.Println("\n ****** variadic parameters  ****** \n ")
	//variadic parameters
	example_variadic("example_variadic", "1", "2", "4")

	fmt.Println("\n ****** optionals can be created using pointers  ****** \n ")
	// optionals can be created using pointers
	example_pointers()

	fmt.Println("\n ****** optionals can be created using structs  ****** \n ")
	// optionals can be created using structs
	example_structs()
}

func example_variadic(requiredParam string, optionalParams ...string) {
	// code
	for _, param := range optionalParams {
		fmt.Printf("printing optionals with variadic %s", param)
	}
}

type Person struct {
	Name string
	Age  *int // Pointer to an int
}

func example_pointers() {
	var age int = 10
	personWithAge := Person{Name: "vru", Age: &age}
	personWithoutAge := Person{Name: "vru", Age: nil}

	printPersonInfo(personWithAge)
	printPersonInfo(personWithoutAge)
}

func printPersonInfo(p Person) {
	fmt.Printf("Name: %s\n", p.Name)
	if p.Age != nil {
		fmt.Printf("Age: %d\n", *p.Age) // Dereference the pointer to get the value
	} else {
		fmt.Println("Age not provided")
	}
	fmt.Println()
}

type OptionalInt struct {
	IsSet bool
	Value int
}

type People struct {
	Name string
	Age  OptionalInt
}

func example_structs() {
	personWithAge := People{Name: "Alice", Age: OptionalInt{IsSet: true, Value: 30}}
	personWithoutAge := People{Name: "Bob", Age: OptionalInt{IsSet: false}}

	printPeopleInfo(personWithAge)
	printPeopleInfo(personWithoutAge)
}

func printPeopleInfo(p People) {
	fmt.Printf("Name: %s\n", p.Name)
	if p.Age.IsSet {
		fmt.Printf("Age: %d\n", p.Age.Value)
	} else {
		fmt.Println("Age not provided")
	}
	fmt.Println()
}
