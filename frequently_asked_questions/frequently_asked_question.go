package main

import (
	"fmt"
	"net"
	"runtime"
)

func main() {
	//concurrency with goroutine
	go func() {
		fmt.Println("Hello From goroutine")
	}()

	//Fast compilation and execution
	fmt.Println("Hello, Go")

	//Cross-platform compatibility
	fmt.Println("Go is awesome on", runtime.GOOS)

	//strong typing
	var num int = 42
	fmt.Println("My favourite number is", num)

	//Built-in support for network applications
	conn, _ := net.Dial("tcp", "example.com:80")
	fmt.Println("Connected to", conn.RemoteAddr())

	// Go routine demo
	func_go_routine_thread()

	//Function vs Method Difference
	fmt.Println("\n ********************* Function vs Method Difference ****************")
	greet("Alice") // Calling the function

	person := Person{Name: "Bob"}
	person.introduce() // Calling the method

	//Threads vs GoRoutines Difference
	fmt.Println("\n ********************* Threads vs GoRoutines Difference ****************")
	func_go_routine_thread()

	fmt.Println("\n ********************* Communicate between multiple goroutines ****************")
	func_multiple_goroutines_with_chan()

	fmt.Println("\n **** Variadic Functions and Usage   ****")
	fmt.Println(sum(1, 2, 3, 4, 5, 6, 7))

	fmt.Println("\n **** Working of Interface in Go  ****")
	interface_demo()

	fmt.Println("\n **** Difference Between Pointer Receiver Method and Value Receiver Method: ****")
	func_pointer_recv_value_recv()

	fmt.Println("\n ***** Use of Select Statement and Switch Statements in Go: **")
	func_select_switch()
}
