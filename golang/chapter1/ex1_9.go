package chapter1

import (
	"fmt"
	"strings"
)

// RunChapter1Exercise9
// String Rotation: Assume you have a method isSubstring which checks if one word is a substring
// of another. Given two strings, s1 and s2, write code to check if s2 is a rotation of s1 using only one
// call to isSubstring (e.g., "waterbottle" is a rotation of "erbottlewat")
func RunChapter1Exercise9() {
	fmt.Println(isRotation("waterbottle", "erbottlewat")) // Expected: true
	fmt.Println(isRotation("hello", "llohe"))             // Expected: true
	fmt.Println(isRotation("hello", "olelh"))             // Expected: false
	fmt.Println(isRotation("abcde", "cdeab"))             // Expected: true
	fmt.Println(isRotation("abcde", "abced"))             // Expected: false
	fmt.Println(isRotation("aaaa", "aaaa"))               // Expected: true
	fmt.Println(isRotation("aab", "aba"))                 // Expected: true
	fmt.Println(isRotation("aab", "baa"))                 // Expected: true
	fmt.Println(isRotation("abc", "cab"))                 // Expected: true
	fmt.Println(isRotation("", ""))                       // Expected: true
	fmt.Println(isRotation("a", ""))                      // Expected: false
	fmt.Println(isRotation("", "a"))                      // Expected: false
}

// Time: O(N)
// Space: O(N)
func isRotation(s1, s2 string) bool {
	// concat the string twice is enought to found all rotations
	duplicated := s1 + s1
	return isSubstring(duplicated, s2)
}

func isSubstring(s1, s2 string) bool {
	return strings.Contains(s1, s2)
}
