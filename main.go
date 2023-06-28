package main

import (
	"strings"
	"os"
)

func main() {
	filename := getFilename()
	text, err := getText(filename)
	if (err != nil) {
		println("Error: ", err)
		os.Exit(1)
	}
	words := strings.Split(filterText(text), " ")
	symbols := proceedWords(words)
	symbol, exists := findFirstUniqueSymbol(symbols)
	if (exists) {
		println(symbol)
	} else {
		println("Such a symbol doesn't exist.")
	}
}

// Returns symbol and true, if unique symbol exists OR empty string and false, if it doesn't
func findFirstUniqueSymbol(s []string) (string, bool) {
	str := strings.Join(s, "")
	for _, c := range str {
		if (strings.Count(str, string(c)) == 1) {
			return string(c), true
		}
	}

	return "", false
}

// Returns slice of first unique symbols in each word
func proceedWords(words []string) []string {
	symbols := make([]string, 0, len(words))
	for _, w := range words {
		if (w == "") {
			continue
		}
		x := ""
		for _, s := range w {
			if (strings.Count(w, string(s)) == 1) {
				x = string(s)
				break
			}
		}
		if (x != "") {
			symbols = append(symbols, x)
		}
	}

	toReturn := make([]string, len(symbols))
	copy(toReturn, symbols)

	return toReturn
}

// Removes all the characters, that are listed in ignored_symbols.txt file, from the given text and returns new string
func filterText(s string) string {
	f, err := os.Open("ignored_symbols.txt")
	defer f.Close()
	if (err != nil) {
		println("Error: ", err)
		os.Exit(1)
	}

	buf := make([]byte, 1024*4) // 4 KB
	c, err := f.Read(buf)
	if (err != nil) {
		println("Error: ", err)
	}
	ignored := string(buf[:c])

	for _, c := range ignored {
		s = strings.ReplaceAll(s, string(c), "")
	}

	return s
}

// Returns filename from console arguments
func getFilename() string {
	filename := "input.txt"
	if len(os.Args) >= 2 {
		filename = os.Args[1]
	}

	return filename
}

// Opens given file and reads the data from it. Returns text and error object
func getText(filename string) (string, error) {
	f, err := os.Open(filename)
	defer f.Close()
	if (err != nil) {
		return "", err
	}
	buf := make([]byte, 1024 * 64) // 64 KB
	c, err := f.Read(buf)
	if (err != nil) {
		return "", err
	}

	return string(buf[:c]), nil
}
