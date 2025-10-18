package chapter1

import (
	"fmt"

	"github.com/ruancaetano/cracking-the-coding-interview/golang/util"
)

var spaceEncodeRunes = []rune{'%', '2', '0'}

// RunChapter1Exercise3
// URLify: Write a method to replace all spaces in a string with '%20'.
// You may assume that the string has sufficient space at the end to hold the additional characters, and that you are given the "true" length of the string.
// (Note: If implementing in Java, please use a character array so that you can perform this operation in place.)
func RunChapter1Exercise3() {
	fmt.Println(urlify(util.StringToRuneArray("123"), 3))
	fmt.Println(urlify(util.StringToRuneArray("1 2 3       "), 6))
	fmt.Println(urlify(util.StringToRuneArray("Mr John Smith    "), 13))

	fmt.Println("----------")

	fmt.Println(urlifyWithBuilder(util.StringToRuneArray("123"), 3))
	fmt.Println(urlifyWithBuilder(util.StringToRuneArray("1 2 3       "), 6))
	fmt.Println(urlifyWithBuilder(util.StringToRuneArray("Mr John Smith    "), 13))
}

// Solution without new structures
// Time: O(N)
// Space: O(1)
// Where N is the input length
func urlify(input []rune, trueLength int) string {
	var resultPosition = len(input) - 1
	for i := trueLength - 1; i >= 0; i-- {
		if input[i] != ' ' {
			input[resultPosition] = input[i]
			resultPosition--
			continue
		}

		input[resultPosition] = '0'
		input[resultPosition-1] = '2'
		input[resultPosition-2] = '%'
		resultPosition -= 3
	}

	return string(input)
}

// Solution with string builder
// Time: O(N)
// Space: O(N)
// Where N is the input length
func urlifyWithBuilder(input []rune, trueLength int) string {
	builder := util.NewStringBuilder()

	for i := 0; i < trueLength; i++ {
		if input[i] != ' ' {
			builder = builder.InsertAtEnd(input[i])
			continue
		}

		builder = builder.InsertAtEnd(spaceEncodeRunes...)
	}

	return builder.Build()
}
