package arraysandstrings

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
