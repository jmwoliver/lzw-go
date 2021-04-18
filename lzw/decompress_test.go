package lzw

import (
	"testing"

	"gotest.tools/assert"
)

func TestDecompress(t *testing.T) {

	file := "./test_data/compressed.lzw"
	returnedByte, err := Decompress(file)
	assert.NilError(t, err)

	returnedString := string(returnedByte)
	expectedString := "Willard Carroll Smith Jr. (born September 25, 1968) is an American actor, rapper, and film producer. Smith has been nominated for five Golden Globe Awards and two Academy Awards, and has won four Grammy Awards."
	assert.Equal(t, returnedString, expectedString)

}
