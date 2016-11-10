package utils

import (
	"bytes"
)

// Characters set used for encoding
const encodeSet = "abcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"

// Encode a integer and return it's base 62 format
func Encode(number int) string {
	if number < 0 {
		panic("Input number should be bigger than 0")
	}
	if number == 0 {
		return string(encodeSet[0])
	}
	var buffer bytes.Buffer
	length := len(encodeSet)

	for number > 0 {
		tmp := number / length
		rmd := number % length
		buffer.WriteString(string(encodeSet[rmd]))
		number = tmp
	}
	return buffer.String()
}
