package main

import "fmt"

func main() {
	builtintypes()

	variables()

	constants()

	fmt.Println("\n **** LoopsDemo ******* \n")
	loopsDemo()

	fmt.Println("\n \b*****  SwitchDemo ***** \n")
	switchDemo()

	fmt.Println("\n \b*****  Arrays and slices Demo ***** \n")
	array_slicesDemo()

	fmt.Println("**********  Maps *************")
	mapsDemo()

	fmt.Println("****** Struct Demo ********")
	structsDemo()

	fmt.Println("**** Ranges ******")
	ranges()

	fmt.Println("**** Pointers ******")
	pointers()

	fmt.Println("****** Interfaces **********")
	interfaces()

	fmt.Println("***** Functions *****")
	functionsDemo()

	fmt.Println("*** GoRoutines *****")
	go_routine_demo()

	fmt.Println("*** Channel ***** ")
	channelsDemo()
}
