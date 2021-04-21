package main

import "fmt"

func main() {
	i := 67
	// arguments to the function call
	// in the defer statement are 
	// evaluated when defer statements
	// are evaluated
	defer fmt.Println(i)
	fmt.Println(test())
	i += 34
}

func test() int {
	return 45;
}