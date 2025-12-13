package chapter1

import (
	"fmt"

	"github.com/ruancaetano/cracking-the-coding-interview/golang/shared"
)

// RunChapter1Exercise5
// One Away: There are three types of edits that can be performed on strings: insert a character, remove a character, or replace a character.
// Given two strings, write a function to check if they are one edit (or zero edits) away.
// EXAMPLE
// pale, pie -> true
// pales, pale -> true
// pale, bale -> true
// pale, bake -> false
func RunChapter1Exercise5() {
	fmt.Println(isOneWayEdit(shared.StringToRuneArray("pale"), shared.StringToRuneArray("pile")))
	fmt.Println(isOneWayEdit(shared.StringToRuneArray("pales"), shared.StringToRuneArray("pale")))
	fmt.Println(isOneWayEdit(shared.StringToRuneArray("pale"), shared.StringToRuneArray("bale")))
	fmt.Println(isOneWayEdit(shared.StringToRuneArray("pale"), shared.StringToRuneArray("bakerrr")))
}

// Time: O(N + K)
// Space: O(N)
// Where N and K is the inputs length respectively
func isOneWayEdit(first []rune, second []rune) bool {
	mappedFirstLetters := map[rune]int{}

	for _, letter := range first {
		if _, exists := mappedFirstLetters[letter]; !exists {
			mappedFirstLetters[letter] = 1
			continue
		}
		mappedFirstLetters[letter] += 1
	}

	differences := 0

	for i := 0; i < len(second); i++ {
		letter := second[i]
		if _, exists := mappedFirstLetters[letter]; !exists {
			differences += 1
			continue
		}

		mappedFirstLetters[letter] -= 1
		if mappedFirstLetters[letter] == 0 {
			delete(mappedFirstLetters, letter)
		}
	}

	return differences <= 1
}
