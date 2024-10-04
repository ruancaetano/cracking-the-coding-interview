package util

func StringToRuneArray(input string) []rune {
	result := make([]rune, 0, len(input))

	for _, char := range input {
		result = append(result, char)
	}
	return result
}
