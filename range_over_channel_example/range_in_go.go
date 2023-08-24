package main

import "fmt"

// range with Arrays, slice , maps, strings, and channel we just saw above

func range_in_go() {

	// range with array and slice
	// To iterate over an array or slice, we can use the range keyword followed by the array or slice variable.
	//can omit index using _, v
	fmt.Println("*** range with array  *** \n ")
	arr := []int{1, 2, 3, 4, 5}
	for i, v := range arr {
		fmt.Printf("index: %d, value : %d \n", i, v)
	}

	fmt.Println("*** range with maps *** \n ")
	// maps
	maap := map[string]string{
		"name": "Riyu",
		"age":  "1 year",
	}

	for k, v := range maap {
		fmt.Printf("key: %s, value: %s\n", k, v)
	}

	fmt.Println("*** range with strings *** \n ")
	//strings
	s := "hello"

	for _, c := range s {
		fmt.Printf("%c\n", c)
	}
}
