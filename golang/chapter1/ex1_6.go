package chapter1

import (
	"fmt"
	"strconv"

	"github.com/ruancaetano/cracking-the-coding-interview/golang/shared"
)

// RunChapter1Exercise6
// String Compression: Implement a method to perform basic string compression using the counts of repeated characters.
// For example, the string aabcccccaaa would become a2b1c5a3, If the "compressed" string would not become smaller than the
// original string, your method should return the original string.
// You can assume the string has only uppercase and lowercase letters (a - z).
func RunChapter1Exercise6() {
	fmt.Println(compressString(shared.StringToRuneArray("aabcccccaaa")))
	fmt.Println(compressString(shared.StringToRuneArray("abc")))
	fmt.Println(compressString(shared.StringToRuneArray("aabbcccccccd")))
}

// Time: O(N)
// Space: O(N)
// Where N is the inputs length
func compressString(original []rune) string {
	result := make([]rune, 0, len(original))

	lastChar := original[0]
	lastCharCount := 1
	for i := 1; i < len(original); i++ {
		if lastChar == original[i] {
			lastCharCount += 1
			continue
		}

		result = append(result, lastChar)
		result = append(result, []rune(strconv.Itoa(lastCharCount))...)
		lastChar = original[i]
		lastCharCount = 1
	}

	result = append(result, lastChar)
	result = append(result, []rune(strconv.Itoa(lastCharCount))...)

	if len(result) > len(original) {
		return string(original)
	}
	return string(result)
}
