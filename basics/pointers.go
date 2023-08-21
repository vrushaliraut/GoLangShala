package main

import "fmt"

type ninja1 struct {
	name string
}

func pointers() {
	tommy := ninja1{"Tommy"}
	tommyPointer := &tommy
	johnyPointer := &ninja1{"Johnny"}

	var ninjaPointer *ninja1 = new(ninja1)

	fmt.Println("tommyPointer:: ", tommyPointer)
	fmt.Println("johnyPointer:: ", johnyPointer)
	fmt.Println("ninjaPointer:: ", ninjaPointer)
}
