package chapter1

import (
	"fmt"
)

// RunChapter1Exercise2
// Check Permutation: Given two strings, write a method to decide if one is a permutation of the other.
func RunChapter1Exercise2() {
	fmt.Println(AreStringPermutable("aaa", "a"))
	fmt.Println(AreStringPermutable("aaa", "aab"))

	fmt.Println("------")

	fmt.Println(AreStringPermutable("aaa", "aaa"))
	fmt.Println(AreStringPermutable("abc", "cba"))
	fmt.Println(AreStringPermutable("aabbcc", "bcacba"))
}

// Time: O(N)
// Space: O(N)
// Where N is the input length
func AreStringPermutable(first, second string) bool {
	if len(first) != len(second) {
		return false
	}

	mappedFirstLetters := make(map[rune]uint8, len(first))
	for _, letter := range first {
		if _, found := mappedFirstLetters[letter]; !found {
			mappedFirstLetters[letter] = 1
			continue
		}
		mappedFirstLetters[letter] += 1
	}

	for _, letter := range second {
		var letterCount uint8
		var found bool
		if letterCount, found = mappedFirstLetters[letter]; !found {
			return false
		}

		if letterCount == 1 {
			delete(mappedFirstLetters, letter)
			continue
		}

		mappedFirstLetters[letter] -= 1
	}

	return len(mappedFirstLetters) == 0
}
