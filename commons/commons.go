package commons

import (
	"fmt"
	"os"
)

// NewBaseDict gets the base dictionary for compression
func NewBaseDict() map[string]byte {
	return baseDict
}

// NewReverseBaseDict gets the base dictionary for decompression
func NewReverseBaseDict() map[byte]string {

	reverseBaseDict := make(map[byte]string)
	for k, v := range baseDict {
		reverseBaseDict[v] = k
	}
	return reverseBaseDict
}

var baseDict = map[string]byte{

	// Lowercase letters
	"a": 0,
	"b": 1,
	"c": 2,
	"d": 3,
	"e": 4,
	"f": 5,
	"g": 6,
	"h": 7,
	"i": 8,
	"j": 9,
	"k": 10,
	"l": 11,
	"m": 12,
	"n": 13,
	"o": 14,
	"p": 15,
	"q": 16,
	"r": 17,
	"s": 18,
	"t": 19,
	"u": 20,
	"v": 21,
	"w": 22,
	"x": 23,
	"y": 24,
	"z": 25,

	// Uppercase letters
	"A": 26,
	"B": 27,
	"C": 28,
	"D": 29,
	"E": 30,
	"F": 31,
	"G": 32,
	"H": 33,
	"I": 34,
	"J": 35,
	"K": 36,
	"L": 37,
	"M": 38,
	"N": 39,
	"O": 40,
	"P": 41,
	"Q": 42,
	"R": 43,
	"S": 44,
	"T": 45,
	"U": 46,
	"V": 47,
	"W": 48,
	"X": 49,
	"Y": 50,
	"Z": 51,

	// Numbers
	"0": 52,
	"1": 53,
	"2": 54,
	"3": 55,
	"4": 56,
	"5": 57,
	"6": 58,
	"7": 59,
	"8": 60,
	"9": 61,

	// Symbols
	"\\": 62,
	"`":  63,
	"~":  64,
	",":  65,
	".":  66,
	"<":  67,
	">":  68,
	"/":  69,
	"?":  70,
	";":  71,
	":":  72,
	"-":  73,
	"_":  74,
	"'":  75,
	"\"": 76,
	"!":  77,
	"@":  78,
	"#":  79,
	"$":  80,
	"%":  81,
	"^":  82,
	"&":  83,
	"*":  84,
	"(":  85,
	")":  86,
	"{":  87,
	"}":  88,
	"[":  89,
	"]":  90,
	"=":  91,
	"+":  92,
	"|":  93,

	// Spaces, tabs, new lines, and such
	" ":  94,
	"\n": 95,
	"\r": 96,
	"\t": 97,

	/*
		-------------------------------------------------------
		Anything created after 98 will be generated dynamically
		-------------------------------------------------------
	*/

}

func CreateFile(path string, data []byte) error {
	// Create the file at the path specified
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write(data)
	if err != nil {
		return err
	}

	fmt.Printf("File created at: '%s'\n", path)

	return nil
}
