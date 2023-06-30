package search

import (
	"strings"
)

const ONE = 1

// Returns symbol and true, if unique symbol exists OR empty string and false, if it doesn't
func FindFirstUniqueSymbol(s string) (string, bool) {
	for _, c := range s {
		if (strings.Count(s, string(c)) == ONE) {
			return string(c), true
		}
	}

	return "", false
}
