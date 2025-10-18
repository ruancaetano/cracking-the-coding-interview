package chapter1

import (
	"fmt"
	"unicode"

	"github.com/ruancaetano/cracking-the-coding-interview/golang/util"
)

// RunChapter1Exercise4
// Given a string, write a function to check if it is a permutation of a palindrome.
// A palindrome is a word or phrase that is the same forwards and backwards.
// A permutation is a rearrangement of letters.The palindrome does not need to be limited to just dictionary words.
// You can ignore casing and non-letter characters
func RunChapter1Exercise4() {
	fmt.Println(isPermutationOfPalindrome(util.StringToRuneArray("Tact Coa")))
	fmt.Println(isPermutationOfPalindrome(util.StringToRuneArray("Amanapla nacana lPanama")))
	fmt.Println(isPermutationOfPalindrome(util.StringToRuneArray("Ablewas 12 IsawElba")))
}

// Time: O(N)
// Space: O(N)
// Where N is the input length
func isPermutationOfPalindrome(input []rune) bool {
	countedLetters := map[rune]int{}

	for _, letter := range input {
		if !unicode.IsLetter(letter) && !unicode.IsNumber(letter) {
			continue
		}

		letter = unicode.ToLower(letter)
		if _, exists := countedLetters[letter]; !exists {
			countedLetters[letter] = 1
			continue
		}
		countedLetters[letter] += 1
	}

	foundOdd := false
	for _, count := range countedLetters {
		if count%2 != 0 {
			if foundOdd {
				return false
			}
			foundOdd = true
		}
	}

	return true
}
