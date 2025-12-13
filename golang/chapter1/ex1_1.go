package chapter1

import (
	"fmt"
	"sort"

	"github.com/ruancaetano/cracking-the-coding-interview/golang/shared"
)

// RunChapter1Exercise1
// Is Unique: Implement an algorithm to determine if a string has all unique characters.
// What if you cannot use additional data structures?
func RunChapter1Exercise1() {
	fmt.Println(hasUniqueChar("44"))
	fmt.Println(hasUniqueChar("117"))
	fmt.Println(hasUniqueChar("132"))

	fmt.Println("---")

	fmt.Println(hasUniqueCharWithoutHashMap(shared.StringToRuneArray("44")))
	fmt.Println(hasUniqueCharWithoutHashMap(shared.StringToRuneArray("117")))
	fmt.Println(hasUniqueCharWithoutHashMap(shared.StringToRuneArray("132")))
}

// Solution with map
// Time: O(N)
// Space: O(N)
// Where N is the input length
func hasUniqueChar(input string) bool {
	charHistory := make(map[rune]bool, len(input))
	for _, char := range input {
		if _, found := charHistory[char]; found {
			return false
		}
		charHistory[char] = true
	}

	return true
}

// Solution without extra structures
// Time  O(N log N + N) -> O(N LOG N)
// Space O(1)
// Where N is the input length
func hasUniqueCharWithoutHashMap(input []rune) bool {
	sort.Slice(input, func(i, j int) bool {
		return input[i] > input[j]
	})

	for i := range input {
		if i+1 >= len(input) {
			return true
		}

		if input[i] == input[i+1] {
			return false
		}
	}
	return true
}
