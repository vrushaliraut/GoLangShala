package main

import "fmt"

/*

- Using the _, ok idiom:
- Using the val, ok := m[key] syntax:
- Using the for range loop:

*/

func check_map_key() {

	// - Using the _, ok idiom:

	m := map[string]int{
		"apple":  1,
		"banana": 2,
		"orange": 3,
	}

	if _, ok := m["apple"]; ok {
		fmt.Println("Key exists!")
	} else {
		fmt.Println("Key does not exist!")
	}

	//- Using the val, ok := m[key] syntax:
	map1 := map[string]int{
		"apple":  1,
		"banana": 2,
		"orange": 3,
	}
	val, ok := map1["apple"]
	if ok {
		fmt.Println("Key exists!", val)
	} else {
		fmt.Println("Key does not exist!")
	}

	//use for range loop

	maap := map[string]int{
		"apple":  1,
		"banana": 2,
		"orange": 3,
	}

	for key, value := range maap {
		if key == "apple" {
			fmt.Println("Key exists!", value)
			break
		}
	}

}
