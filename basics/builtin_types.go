package main

import "fmt"

/*
    bool
	string
	int, int8, int16, int32, int64
	uint, uint8, uint16, uint32, uint64,uintptr,
	rune, byte
	float32,float64,
	complex64, complex132
*/

func builtintypes() {
	var str string = "ABCD"
	r_array := []rune(str)
	fmt.Println("----BuiltIn types-----")

	fmt.Printf("Array of rune values for A, B, C and D: %U\n", r_array)
}
