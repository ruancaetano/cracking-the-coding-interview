package chapter1

import (
	"fmt"
	"sort"

	"github.com/ruancaetano/cracking-the-coding-interview/golang/util"
)

// Run_1_1 Is Unique: Implement an algorithm to determine if a string has all unique characters.
// What if you cannot use additional data structures?
func Run1_1() {
	fmt.Println(hasUniqueChar("44"))
	fmt.Println(hasUniqueChar("117"))
	fmt.Println(hasUniqueChar("132"))

	fmt.Println("---")

	fmt.Println(hasUniqueCharWithoutHashMap("44"))
	fmt.Println(hasUniqueCharWithoutHashMap("117"))
	fmt.Println(hasUniqueCharWithoutHashMap("132"))
}

// O(n)
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

// O( N + N log N + N)  -> O(N LOG N)
func hasUniqueCharWithoutHashMap(input string) bool {
	inputAsArray := util.StringToRuneArray(input)

	sort.Slice(inputAsArray, func(i, j int) bool {
		return inputAsArray[i] > inputAsArray[j]
	})

	for i, _ := range inputAsArray {
		if i+1 >= len(inputAsArray) {
			return true
		}

		if inputAsArray[i] == inputAsArray[i+1] {
			return false
		}
	}
	return true
}
