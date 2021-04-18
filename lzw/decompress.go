package lzw

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/jmwoliver/lzw-go/commons"
)

func Decompress(file string) ([]byte, error) {

	// Create the reverse base dictionary with all characters on US keyboard
	dict := commons.NewReverseBaseDict()
	count := len(dict)

	// Save decompressed data to a string
	var decompressedString string

	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	r := bufio.NewReader(f)
	for {
		b, err := r.ReadByte()
		if err == io.EOF {
			break
		} else if err != nil {
			// Error reading the file, bail out because it busted
			return nil, err
		} else {
			// Can guarentee this will be in the dictionary based on how LZW builds it
			dictVal := dict[b]

			// Start generating next key, value pair for dictionary
			dictKey := byte(count)
			dict[dictKey] = dictVal

			// Put the value in the decompressed string
			decompressedString += dictVal

			// Get the next byte to create the next dictionary entry
			nb, err := r.ReadByte()
			if err == io.EOF {
				break
			} else if err != nil {
				// Error reading the file, bail out because it busted
				return nil, err
			} else {
				nextVal := dict[nb]
				nextDictVal := dictVal + string(nextVal[0])
				dict[dictKey] = nextDictVal

				// Unread the last byte to continue reading through the file at the correct location
				r.UnreadByte()
				count += 1
			}
		}
	}

	fmt.Printf("Decompressed string looks like: %s\n", decompressedString)
	return []byte(decompressedString), nil
}
