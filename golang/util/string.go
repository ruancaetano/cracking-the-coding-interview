package util

// StringToRuneArray converts a string to a mutable []rune slice for algorithm exercises.
// This utility allows to work with mutable character arrays while keeping test inputs readable.
//
// Examples:
//
//	StringToRuneArray("hello")     // returns []rune{'h','e','l','l','o'}
//	StringToRuneArray("Mr John")   // returns []rune{'M','r',' ','J','o','h','n'}
func StringToRuneArray(input string) []rune {
	result := make([]rune, 0, len(input))

	for _, char := range input {
		result = append(result, char)
	}
	return result
}
