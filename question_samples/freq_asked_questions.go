package main

import "fmt"

func main() {

	// manage multiline strings
	func_multiline()

	//enums with goorm
	enums_with_goorm()

	//map_equality
	map_equality()

	//slices_equality
	slices_equality()

	//optionals
	explore_different_optionals()

	//check map key efficiently
	check_map_key()
}

func func_multiline() {
	//raw string literals
	message := `Hello 
	World`
	fmt.Println(message)

	//concatination
	message1 := "Hello " +
		"World"
	fmt.Println(message1)

	//using fmt.sprintf
	message3 := fmt.Sprintf("Hello\nWorld")
	fmt.Println(message3)
}
