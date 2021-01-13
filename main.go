package main

import (
	"fmt"

	"github.com/benipal95/cracking-go/arraysandstrings"
)

func main() {
	stringData := "A quick brown fox jumps over a lazy dog"
	string1 := "a quick"
	string2 := "quick a"
	urlString := "Mr John Smith    "
	permPalindrome := "race car"
	string3 := "pas"
	string4 := "pale"
	/* Arrays and Strings coding problems */
	fmt.Printf("Q: 1.1 - Input: %v - IsUnique: %v \n", stringData, arraysandstrings.IsUnique(stringData))
	fmt.Printf("Q: 1.2 - Input: %v, %v - isPermutation: %v \n", string1, string2, arraysandstrings.CheckPermutation(string1, string2))
	fmt.Printf("Q: 1.3 - Input: %v - URLify: %v \n", urlString, arraysandstrings.URLify(urlString, 13))
	fmt.Printf("Q: 1.4 - Input: %v - PalindromePermutation: %v \n", permPalindrome, arraysandstrings.PalindromePermutation(permPalindrome))
	fmt.Printf("Q: 1.5 - Input: %v, %v - OneAway: %v \n", string3, string4, arraysandstrings.OneAway(string3, string4))

}
