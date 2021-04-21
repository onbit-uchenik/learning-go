/*
	Author - Anubhav Vats
	Description - Calculate the sqrt of any positive number.
	Last Modified - 21/04/21
*/
package main

import (
	"errors"
	"fmt"
	"math"
	"reflect" // for getting type of any value.
)

func Sqrt(x float64) (error, float64) {
	fmt.Println("checking the type of x", reflect.TypeOf(x)) // prints float64
	if x < 0 {
		return errors.New("expected input to Sqrt to be positive number, found negative number"), 0
	}
	z := x / 2 // assumed root
	limit := 0.00000000000001
	count := 0
	for root := z - ((z*z)-x)/(2*z); math.Abs(root-z) > limit; {
		z = root
		root = z - ((z*z)-x)/(2*z)
		count++
	}
	return nil, z
}

func main() {
	// There is no implicit type conversion in go lang
	// still how sending 45 a int value to float64 didnt throw any
	// error?
	fmt.Println("Checking type of the constant", reflect.TypeOf(45)) // prints int
	err, sqNumber := Sqrt(45)
	if err == nil {
		fmt.Println("Square root of 45 is ", sqNumber)
	} else {
		fmt.Println(err)
	}
}
