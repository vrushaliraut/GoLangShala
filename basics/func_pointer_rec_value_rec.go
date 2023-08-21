package main

import "fmt"

type Counter struct {
	Count int
}

// Pointer receiver method
func (c *Counter) Increment() {
	c.Count++
}

// Value Receiver Method
func (c Counter) GetValue() int {
	return c.Count
}

func func_pointer_recv_value_recv() {
	c := Counter{}
	c.Increment()             // Pointer receiver method
	fmt.Println(c.GetValue()) // Value receiver method
}
