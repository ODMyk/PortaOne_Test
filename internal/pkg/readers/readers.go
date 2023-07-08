package readers

import (
	"os"
)

const (
	INPUT_FILENAME = "input.txt"
	IGNORED_SYMBOLS_FILENAME = "ignored_symbols.txt"
	TWO = 2
)

// Returns filename from console arguments
func GetFilenames() (string, string) {
	input := INPUT_FILENAME
	ignore := IGNORED_SYMBOLS_FILENAME
	if len(os.Args) >= TWO {
		input = os.Args[1]
		if (len(os.Args) > TWO) {
			ignore = os.Args[2]
		}
	}

	return input, ignore
}

// Opens given file and reads the data from it. Returns text and error object
func GetText(filename string) (string, error) {
	f, err := os.Open(filename)
	defer f.Close()
	if (err != nil) {
		return "", err
	}
	fInfo, err := f.Stat()
	if (err != nil) {
		return "", err
	}
	buf := make([]byte, fInfo.Size())
	c, err := f.Read(buf)
	if (err != nil) {
		return "", err
	}

	return string(buf[:c]), nil
}
