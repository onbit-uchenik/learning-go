package main // in Go, main package is must for an application to run

import (
	"fmt"
	"example.com/fibonacci"
	"time"
)

func main() {
	start := time.Now()
	num := uint64(1 << 63)
	ans := fibonacci.FiboLessThanN(num)
	duration := time.Since(start).Nanoseconds()
	fmt.Printf("fibonacci number less than %d is %d\n", num, ans)
	fmt.Printf("Time Taken := %d nanoseconds\n", duration)	
}