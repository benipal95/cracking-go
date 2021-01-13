package main

import (
	"fmt"

	"github.com/benipal95/cracking-go/arraysandstrings"
)

func main() {
	stringData := "A quick brown fox jumps over a lazy dog"
	string1 := "a quick"
	string2 := "quick a"

	/* Arrays and Strings coding problems */
	fmt.Printf("Q: 1.1 - Data: %v - IsUnique: %v \n", stringData, arraysandstrings.IsUnique(stringData))
	fmt.Printf("Q: 1.2 - Data: %v - isPermutation: %v \n", stringData, arraysandstrings.CheckPermutation(string1, string2))
}
