package PortaOne_Test

import (
	"log"
	"strings"
	"PortaOne_Test/internal/pkg/readers"
	"PortaOne_Test/pkg/search"
)

func Run() {
	inputName, filterName := readers.GetFilenames()
	text, err := readers.GetText(inputName)
	if (err != nil) {
		log.Fatal(err)
	}
	filter, err := readers.GetText(filterName)
	if (err != nil) {
		log.Fatal(err)
	}
	words := strings.Split(filterText(text, filter), " ")
	symbols := proceedWords(words)
	symbol, exists := search.FindFirstUniqueSymbol(strings.Join(symbols, ""))
	s := "Such a symbol doesn't exist."
	if (exists) {
		s = "'" + symbol + "'"
	}
	log.Println(s)
}

// Returns slice of first unique symbols in each word
func proceedWords(words []string) []string {
	symbols := make([]string, 0, len(words))
	for _, w := range words {
		s, exists := search.FindFirstUniqueSymbol(w)
		if (exists) {
			symbols = append(symbols, s)
		}
	}

	toReturn := make([]string, len(symbols))
	copy(toReturn, symbols)

	return toReturn
}

// Removes all the characters, that are listed in ignored_symbols.txt file, from the given text and returns new string
func filterText(s, filter string) string {
	for _, c := range filter {
		s = strings.ReplaceAll(s, string(c), "")
	}

	return s
}
