package flags

import (
	"errors"
	"flag"
	"path"
)

func GetFlags() (string, bool, error) {

	// TODO print usage if no flags are passed in or if --help is used

	// Get values from flags
	file, compress, decompress := parseFlags()

	// Validation of flags
	err := validateUsage(file, compress, decompress)
	if err != nil {
		return "", false, err
	}

	// Make sure using right filetype for compress or decompress
	err = validateFiletype(file, compress, decompress)
	if err != nil {
		return "", false, err
	}

	// Boolean value table:
	// true -> compress
	// false -> decompress
	if compress {
		return file, true, nil
	} else {
		return file, false, nil
	}
}

func parseFlags() (string, bool, bool) {

	// Declare file, compress, and decompress variables
	var file string
	var compress bool
	var decompress bool

	// Declare flags and their usage
	flag.StringVar(&file, "file", "", "Specifies the file")
	flag.BoolVar(&compress, "compress", false, "Specifies if compressing a file")
	flag.BoolVar(&decompress, "decompress", false, "Specifies if decompressing a file")

	// Now parse flags to use them
	flag.Parse()

	return file, compress, decompress
}

func validateUsage(file string, compress, decompress bool) error {
	// Must have a file path set
	if file == "" {
		return errors.New("Must have a file path set")
	}

	// Can only have compress or decompress, not both
	if !compress && !decompress {
		return errors.New("Must set either compress or decompress")
	}

	// Can only have compress or decompress, not both
	if compress && decompress {
		return errors.New("Can only use compress or decompress, not both")
	}

	return nil
}

func validateFiletype(file string, compress, decompress bool) error {
	_, fileName := path.Split(file)
	fileType := path.Ext(fileName)

	// Check if no filetype
	if fileType == "" {
		return errors.New("A filetype is needed")
	}

	// check if correct filetype is used if compress flag is enabled
	if compress && fileType != ".txt" {
		return errors.New("The compress flag must be used with a .txt file")
	}

	// check if correct filetype is used if decompress flag is enabled
	if decompress && fileType != ".lzw" {
		return errors.New("The compress flag must be used with an .lzw file")
	}

	return nil
}
