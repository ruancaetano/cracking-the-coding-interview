package chapter1

import (
	"fmt"
	"sort"
	"strings"
	"unicode"

	"github.com/ruancaetano/cracking-the-coding-interview/golang/util"
)

// Run1_4 Given a string, write a function to check if it is a permutation of a palindrome.
// A palindrome is a word or phrase that is the same forwards and backwards.
// A permutation is a rearrangement of letters.The palindrome does not need to be limited to just dictionary words.
// You can ignore casing and non-letter characters
func Run1_4() {
	fmt.Println(isPermutationOfPalindrome(util.StringToRuneArray("Tact Coa")))
	fmt.Println(isPermutationOfPalindrome(util.StringToRuneArray("Amanapla nacana lPanama")))
	fmt.Println(isPermutationOfPalindrome(util.StringToRuneArray("Ablewas 12 IsawElba")))

	fmt.Println("-----------")

	fmt.Println(generatePermutationFromPalindrome(util.StringToRuneArray("Tact Coa")))
	fmt.Println(generatePermutationFromPalindrome(util.StringToRuneArray("Amanapla nacana lPanama")))
	fmt.Println(generatePermutationFromPalindrome(util.StringToRuneArray("Ablewas 12 IsawElba")))
}

// Time and Space O(N)
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

// Time (N + K LOG K) -> O(N) or O(N LOG N) if K is near of N (worse case)
// Space O(N)
// Where N string size, K unique chars and N >= K
func generatePermutationFromPalindrome(input []rune) string {
	countedLetters := map[rune]int{}
	var lettersRankedByCount []rune

	builder := util.NewStringBuilder()
	for _, letter := range input {
		if !unicode.IsLetter(letter) && !unicode.IsNumber(letter) {
			continue
		}

		letter = unicode.ToLower(letter)
		if _, exists := countedLetters[letter]; !exists {
			countedLetters[letter] = 1
			lettersRankedByCount = append(lettersRankedByCount, letter)
			continue
		}
		countedLetters[letter] += 1
	}

	sort.Slice(lettersRankedByCount, func(i, j int) bool {
		letterA := lettersRankedByCount[i]
		letterB := lettersRankedByCount[j]
		return countedLetters[letterA] < countedLetters[letterB]
	})

	for _, letter := range lettersRankedByCount {
		count := countedLetters[letter]
		if count == 1 {
			builder = builder.InsertAtStart(string(letter))
			continue
		}

		toAddOnStart := count / 2
		toAddOnEnd := count - toAddOnStart
		if toAddOnEnd < 0 {
			toAddOnEnd *= -1
		}

		builder = builder.InsertAtStart(strings.Repeat(string(letter), toAddOnStart))
		builder = builder.InsertAtEnd(strings.Repeat(string(letter), toAddOnEnd))
	}

	result := builder.Build()

	resultLen := len(result)
	for i := 0; i < resultLen; i++ {
		if result[i] != result[resultLen-i-1] {
			return "is not palindrome"
		}
	}
	return result
}
