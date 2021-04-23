/*
	Author - Anubhav Vats
	Description - Calculate frequency of each word in a sentence.
	Last Modified - 23/04/21
*/
package main

import (
	"strings"
	"unicode"
	"fmt"
)

// removes punctuation that are presents either in starting or ending
func removePunctuation(word string) string {
	return strings.TrimFunc(word, func(r rune) bool {
		return !unicode.IsLetter(r);
	})
}


func countWords(s string) map[string]int {
	countMap := make(map[string]int)
	substr := strings.Split(s, " ")
	for i := range substr {
		value := removePunctuation(substr[i])
		countMap[value] = countMap[value] + 1
	}

	return countMap
}

func main() {
	myMap := countWords("I I I I I we we we we hello we we we we we we no ni no    ")
	
	fmt.Printf("Number of spaces in above line are %v\n", myMap[""])
}