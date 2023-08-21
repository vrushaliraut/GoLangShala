package main

import "fmt"

func array_slicesDemo() {

	//An array is numbered sequence of elements of specific length
	arrays()

	//A slice on the other hand, doesn't need to be given a specific length
	slices()

}

func arrays() {
	var evilNinjas [3]string
	fmt.Println(len(evilNinjas))

	evilNinjas[0] = "Riyu"
	fmt.Println(evilNinjas)
	fmt.Println(evilNinjas[0])
	fmt.Println(len(evilNinjas))

	moreEvilNinjas := [3]string{"John", "Cator", "Edyo"}
	fmt.Println(moreEvilNinjas)

	var missionReward [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			missionReward[i][j] = i + j
		}
	}

	fmt.Println(missionReward)
}

func slices() {
	fmt.Println("********** slices ********")
	var evilNinja []string
	fmt.Println(len(evilNinja))

	evilNinja = append(evilNinja, "John")
	fmt.Println(len(evilNinja))
}
