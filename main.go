package main

import (
	"log"

	"github.com/jmwoliver/lzw-go/commons"
	"github.com/jmwoliver/lzw-go/flags"
	"github.com/jmwoliver/lzw-go/lzw"
)

func main() {

	file, compress, err := flags.GetFlags()
	if err != nil {
		log.Fatal(err)
	}

	// Run compress or decompress depending on what flag is used
	err = runLzw(file, compress)
	if err != nil {
		log.Fatal(err)
	}
}

func runLzw(file string, compress bool) error {
	// true -> compress
	// false -> decompress
	if compress {
		compressedSlice, err := lzw.Compress(file)
		if err != nil {
			return err
		}
		// Create the compressed file at the current directory
		err = commons.CreateFile("./output/compressed.lzw", compressedSlice)
		if err != nil {
			return err
		}
	} else {
		decompressedSlice, err := lzw.Decompress(file)
		if err != nil {
			return err
		}
		// Create the decompressed file at the current directory
		err = commons.CreateFile("./output/decompressed.txt", decompressedSlice)
		if err != nil {
			return err
		}
	}
	return nil
}
