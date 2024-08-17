package counter

import (
	"strings"
	"unicode"
)

type wordCounter struct {
}

func Word() *wordCounter {
	return &wordCounter{}
}

func (w *wordCounter) Count(text []byte) int {
	trimmed := strings.TrimSpace(string(text))
	count := 0
	isSpace := true
	for _, char := range trimmed {
		if unicode.IsSpace(char) {
			isSpace = true
		} else {
			if isSpace {
				count++
			}
			isSpace = false
		}
	}
	return count
}
