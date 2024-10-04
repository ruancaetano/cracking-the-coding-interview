package chapter1

import (
	"fmt"

	"github.com/ruancaetano/cracking-the-coding-interview/golang/util"
)

const SpaceEncodeString = "%20"

// Run1_3 URLify: Write a method to replace all spaces in a string with '%20'. You may assume that the string has sufficient space at the end to hold the additional characters, and that you are given the "true" length of the string.
// (Note: If implementing in Java, please use a character array so that you can perform this operation in place.)
func Run1_3() {
	fmt.Println(urlify(util.StringToRuneArray("123"), 3))
	fmt.Println(urlify(util.StringToRuneArray("1 2 3       "), 6))
	fmt.Println(urlify(util.StringToRuneArray("Mr John Smith    "), 13))

	fmt.Println("----------")

	fmt.Println(urlifyWithBuilder(util.StringToRuneArray("123"), 3))
	fmt.Println(urlifyWithBuilder(util.StringToRuneArray("1 2 3       "), 6))
	fmt.Println(urlifyWithBuilder(util.StringToRuneArray("Mr John Smith    "), 13))
}

// time O(N), space O(1)
func urlify(input []rune, trueLength int) string {
	spaceCount := 0
	for i := 0; i < trueLength; i++ {
		if input[i] == ' ' {
			spaceCount += 1
			continue
		}
	}

	// empty char will be change to %20, where % replace ` ` and the chars 2 and 0 are added, so each space add 2 new chars)
	resultStringIdx := (trueLength + (spaceCount * 2)) - 1
	for i := trueLength - 1; i >= 0; i-- {
		if input[i] == ' ' {
			input[resultStringIdx] = '0'
			input[resultStringIdx-1] = '2'
			input[resultStringIdx-2] = '%'
			resultStringIdx -= 3
			continue
		}

		input[resultStringIdx] = input[i]
		resultStringIdx -= 1
	}

	return string(input)
}

// time O(N), space O(N)
func urlifyWithBuilder(input []rune, trueLength int) string {
	builder := util.NewStringBuilder()

	for i := 0; i < trueLength; i++ {
		char := string(input[i])
		if char != " " {
			builder = builder.Add(char)
			continue
		}

		builder = builder.Add(SpaceEncodeString)
	}

	return builder.Build()
}
