package main

import (
	"fmt"
	"reflect"
)

func slices_equality() {
	//loop through or deep equality

	slice1 := []int{1, 2, 3}
	slice2 := []int{1, 2, 3}

	if len(slice1) != len(slice2) {
		fmt.Println("Slices are not equal")
		return
	}

	for i, v := range slice1 {
		if v != slice2[i] {
			fmt.Println("Slices are not equal")
			return
		}
	}

	fmt.Println("Slices are equal")

	fmt.Println("\n compairing with deep equal\n")

	slice3 := []int{1, 2, 1}
	slice4 := []int{1, 2, 0}

	if reflect.DeepEqual(slice3, slice4) {
		fmt.Println("Slices are equal")
	} else {
		fmt.Println("Slices are not equal")
	}
}
