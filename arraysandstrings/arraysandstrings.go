package arraysandstrings

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math"
	"strings"
)

// IsUnique determines if a string has all unique characters
func IsUnique(stringData string) bool {
	var hashMap = make(map[rune]bool)
	for _, c := range stringData {
		//Ignore space code point (rune)
		if hashMap[c] && c != 32 {
			return false
		} else {
			hashMap[c] = true
		}
	}
	return true
}

// CheckPermutation returns true if a string is permutation of the other
func CheckPermutation(string1, string2 string) bool {
	if len(string1) != len(string2) {
		return false
	}
	var asciiArray [128]int
	for i := 0; i < len(string1); i++ {
		asciiArray[string1[i]]++
	}
	for i := 0; i < len(string2); i++ {
		asciiArray[string2[i]]--
		if asciiArray[string2[i]] < 0 {
			return false
		}
	}
	return true
}

// URLify replaces space with %20 in a string
// Easier to start from the end during string/array maniplation
func URLify(input string, trueLength int) string {
	inputRune := []rune(input)
	var numOfSpaces = numberOfSpaces(input, 0, trueLength)
	newIndex := trueLength - 1 + numOfSpaces*2

	for oldIndex := trueLength - 1; oldIndex >= 0; oldIndex-- {
		if inputRune[oldIndex] == ' ' {
			inputRune[newIndex] = '0'
			inputRune[newIndex-1] = '2'
			inputRune[newIndex-2] = '%'
			newIndex -= 3
		} else {
			inputRune[newIndex] = inputRune[oldIndex]
			newIndex--
		}
	}
	return string(inputRune)
}

// Get the number of spaces in a string from start to end index
// TODO: Convert to countOfChar()
func numberOfSpaces(input string, start int, end int) (count int) {
	for i := start; i < end; i++ {
		if input[i] == 32 {
			count++
		}
	}
	return
}

// PalindromePermutation returns true if string can be permutated in such a way that it's a palindrome
func PalindromePermutation(input string) bool {
	input = strings.ReplaceAll(input, " ", "")
	str := []rune(input)
	countMap := make(map[rune]int)
	for i := 0; i < len(str); i++ {
		if str[i] != 32 {
			countMap[str[i]] = countMap[str[i]] + 1
		}
	}

	numOfOddKeys := 0
	for _, v := range countMap {
		if v%2 != 0 {
			numOfOddKeys++
		}
	}
	if len(input)%2 == 0 && numOfOddKeys == 0 {
		return true
	} else if len(input)%2 != 0 && numOfOddKeys == 1 {
		return true
	} else {
		return false
	}
}

// OneAway returns true if one of the two strings is one edit away from being equal
func OneAway(str1, str2 string) bool {
	if math.Abs(float64(len(str1)-len(str2))) > 1 {
		return false
	}
	smallString := str1
	bigString := str2
	if len(str1) > len(str2) {
		smallString, bigString = str2, str1
	}

	foundDiff := false
	for i, j := 0, 0; i < len(smallString) && j < len(bigString); {
		if smallString[i] != bigString[j] {
			// Can only have one different char
			if foundDiff {
				return false
			}
			foundDiff = true
			if len(str1) == len(str2) { // move short pointer on replace
				i++
			}
		} else {
			i++
		}
		j++
	}
	return true
}

// StringCompression returns compressed string aaabbc turns into a3b2c
func StringCompression(input string) string {
	var str = []rune(input)
	fmt.Println(str)
	// out := make([]rune, len(str))
	for i := len(input) - 1; i > 0; i-- {
		// curr := str[i]
		// prev := str[i-1]
		count := 1
		for j := i - 1; str[j] == str[i] && j >= 0; j-- {
			count++
			if j == 0 {
				break
			}
		}
		if count > 1 {
			buf := new(bytes.Buffer)
			_ = binary.Write(buf, binary.LittleEndian, count)
			// str[i-count+2] = buf.Bytes()
			var b = []rune{rune(count)}
			str = append(str[:i-count+2], append(b, str[i+1:]...)...)
			fmt.Printf("str: %v, b: %v, index: %v \n", str, b, i-count+2)
			// str = append(str[:i+1], append(b, str[i+count:]...)...)

		}
		i = i - count + 1
	}
	var sb strings.Builder
	for _, v := range str {
		if v == 4 {
			sb.Write([]byte{byte(v)})
		}
	}
	fmt.Println(str)
	// s1 := string(sb.String())
	return sb.String()
}
