package main

import (
	"fmt"
	"reflect"
)

func map_equality() {

	//Iterate through the maps to check equality
	map1 := map[string]int{
		"Abhi":     32,
		"Vrushali": 31,
		"Riyanshi": 1,
	}

	map2 := map[string]int{
		"Abhi":     32,
		"Vrushali": 31,
		"Riyanshi": 2,
	}

	result := maps_equal(map1, map2)
	if result {
		fmt.Println("map are equal")
	} else {
		fmt.Println("map are NOT equal")
	}

	//using refelct.DeepEqual
	result1 := reflect.DeepEqual(map1, map2)
	if result1 {
		fmt.Println("map are equal")
	} else {
		fmt.Println("map are NOT equal")
	}
}

func maps_equal(map1 map[string]int, map2 map[string]int) bool {
	if len(map1) != len(map2) {
		return false
	}

	for k, v1 := range map1 {
		v2, ok := map2[k]
		if !ok || v1 != v2 {
			return false
		}
	}
	return true
}
