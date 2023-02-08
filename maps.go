package main

import "fmt"

func mapsDemo() {

	ninjasLevel := make(map[string]int)
	ninjasLevel["riyu"] = 1
	ninjasLevel["abhi"] = 2

	fmt.Println("print map ::  ", ninjasLevel)
	fmt.Println("Length of map :: ", len(ninjasLevel))

	delete(ninjasLevel, "riyu")
	fmt.Println("Length of map :: ", len(ninjasLevel))

	// another level of initialising maps
	moreninjas := map[string]int{"riyu": 1, "abhi": 2}
	fmt.Println("Length of map with initialise then and there:: ", len(moreninjas))

	//the optional return value if key is present in map
	_, ok := ninjasLevel["abhi"]
	fmt.Println("If key is present ::", ok)
}
