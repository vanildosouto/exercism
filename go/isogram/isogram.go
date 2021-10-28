package isogram

import (
	"strings"
)

//IsIsogram return if word is a isogram
func IsIsogram(word string) bool {
	word = strings.ToLower(word)

	for x := 0; x < len(word); x++ {
		for y := x + 1; y < len(word); y++ {
			if word[x] == 45 || word[x] == 32 {
				continue
			}

			if word[x] == word[y] {
				return false
			}
		}
	}

	return true
}
