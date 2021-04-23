package main

import (
	"fmt"
	"github.com/onbit-syn/learning-go/tree/main/introduction/mypackages"
	"reflect"
)

func main() {
	// There is no implicit type conversion in go lang
	// still how sending 45 a int value to float64 didnt throw any
	// error?
	fmt.Println("Checking type of the constant", reflect.TypeOf(45)) // prints int
	err, sqNumber := mathalgo.Sqrt(45)
	
	if err == nil {
		fmt.Println("Square root of 45 is ", sqNumber)
	} else {
		fmt.Println(err)
	}
}