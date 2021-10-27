package scrabble

import (
	"strings"
)

// Score return a score of word
func Score(word string) (score int) {
	letterValues := map[int][]string{
		1:  {"A", "E", "I", "O", "U", "L", "N", "R", "S", "T"},
		2:  {"D", "G"},
		3:  {"B", "C", "M", "P"},
		4:  {"F", "H", "V", "W", "Y"},
		5:  {"K"},
		8:  {"J", "X"},
		10: {"Q", "Z"},
	}

	word = strings.ToUpper(word)
	for _, letter := range word {
		for value, lettersMap := range letterValues {
			for _, letterInMap := range lettersMap {
				if letterInMap == string(letter) {
					score += value
				}
			}
		}
	}

	return
}
