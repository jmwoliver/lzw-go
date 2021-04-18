package lzw

import (
	"testing"

	"gotest.tools/assert"
)

func TestCompress(t *testing.T) {

	file := "./test_data/original.txt"
	returnedByte, err := Compress(file)
	assert.NilError(t, err)

	returnedString := string(returnedByte)

	expectedByte := []byte{48, 8, 11, 11, 0, 17, 3, 94, 28, 102, 17, 14, 100, 94, 44, 12, 8, 19, 7, 94, 35, 17, 66, 94, 85, 1, 14, 17, 13, 111, 4, 15, 19, 4, 12, 1, 4, 17, 94, 54, 57, 65, 94, 53, 61, 58, 60, 86, 94, 8, 18, 94, 0, 126, 26, 12, 134, 8, 2, 150, 149, 2, 19, 124, 139, 17, 0, 15, 15, 134, 139, 150, 104, 5, 99, 12, 94, 15, 108, 3, 20, 2, 134, 120, 112, 114, 116, 7, 0, 148, 133, 4, 126, 13, 14, 113, 13, 0, 130, 170, 124, 94, 171, 21, 4, 94, 32, 109, 3, 189, 203, 11, 14, 133, 94, 26, 22, 102, 3, 148, 169, 94, 19, 22, 14, 212, 156, 206, 12, 24, 212, 214, 103, 18, 168, 13, 104, 185, 148, 221, 126, 5, 14, 20, 135, 32, 163, 12, 226, 228, 215, 18, 66}
	expectedString := string(expectedByte)
	assert.Equal(t, returnedString, expectedString)
}
