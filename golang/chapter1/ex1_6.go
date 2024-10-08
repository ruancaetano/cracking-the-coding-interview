package chapter1

import (
	"fmt"
	"strconv"

	"github.com/ruancaetano/cracking-the-coding-interview/golang/util"
)

// String Compression: Implement a method to perform basic string compression using the counts of repeated characters.
// For example, the string aabcccccaaa would become a2b1c5a3, If the "compressed" string would not become smaller than the
// original string, your method should return the original string.
// You can assume the string has only uppercase and lowercase letters (a - z).
func Run1_6() {
	fmt.Println(compressString(util.StringToRuneArray("aabcccccaaa")))
	fmt.Println(compressString(util.StringToRuneArray("abc")))
	fmt.Println(compressString(util.StringToRuneArray("aabbcccccccd")))
}

func compressString(original []rune) string {
	var result string

	lastChar := original[0]
	lastCharCount := 1
	for i := 1; i < len(original); i++ {
		if lastChar == original[i] {
			lastCharCount += 1
			continue
		}

		parsedNumber := strconv.Itoa(lastCharCount)
		result += string(lastChar) + parsedNumber
		lastChar = original[i]
		lastCharCount = 1
	}

	parsedNumber := strconv.Itoa(lastCharCount)
	result += string(lastChar) + parsedNumber

	if len(result) > len(original) {
		return string(original)
	}
	return result
}
