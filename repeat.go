package iteration

import "strings"

// Repeat returns a new string consisting of count copies of the rune character
func Repeat(character rune, count int) string {
	var repeatedBuilder strings.Builder
	for i := 0; i < count; i++ {
		repeatedBuilder.WriteRune(character)
	}
	return repeatedBuilder.String()
}
