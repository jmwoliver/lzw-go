package lzw

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/jmwoliver/lzw-go/commons"
)

/*
	LZW Compression High Level Steps:
	1. Instantiate map of all possible single-character values in dictionary
	2. Go through each character 1 at a time and see if they are in the map
	   - Check if n+1 is in the map -> if so, check n+2, etc.
	   - If not -> save previous value to the file, add new key to the map,
	               and roll back the buffer 1 byte to continue compressing.
	3. Return a compressed file with byte values (return an .lzw file)
*/
func Compress(file string) ([]byte, error) {

	// Create the base dictionary with all characters on US keyboard
	dict := commons.NewBaseDict()

	// Save compressed data to a byte slice
	var compressedSlice []byte

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
			return nil, err
		} else {
			// Used to build up the next key in the dictionary
			prevKey := string(b)
			currKey := string(b)
			for {
				if _, contains := dict[currKey]; contains {
					prevKey = currKey
					nb, err := r.ReadByte()
					if err == io.EOF {
						dictVal := dict[currKey]
						compressedSlice = append(compressedSlice, dictVal)
						break
					} else if err != nil {
						log.Fatal("Error reading file: ", err)
						break
					} else {
						currKey += string(nb)
					}
				} else {
					// If currKey isn't contained in the dictionary, then set the dictVal to the prevKey which is the last successful key
					dictVal := dict[prevKey]

					// Append the dictionary value to the compressed slice
					compressedSlice = append(compressedSlice, dictVal)

					// Add next key that wasn't already in the map at the end of the dictionary
					dict[currKey] = byte(len(dict))

					// Unread the last byte to continue reading through the file at the correct location
					r.UnreadByte()
					break
				}
			}
		}
	}

	fmt.Printf("Compressed byte array looks like: %v\n", compressedSlice)
	return compressedSlice, nil
}
