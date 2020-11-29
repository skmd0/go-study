package main

import (
	"bytes"
	"fmt"
)

// ---------------------------------------------------------
// EXERCISE: Append
//
//  Please follow the instructions within the code below.
//
// EXPECTED OUTPUT
//  They are equal.
//
// HINTS
//  bytes.Equal function allows you to compare two byte
//  slices easily. Check its documentation: go doc bytes.Equal
// ---------------------------------------------------------

func main() {
	var header []byte
	png := []byte{'P', 'N', 'G'}
	header = append(header, png...)

	areEqual := bytes.Equal(png, header)
	if areEqual {
		fmt.Println("They are equal.")
	} else {
		fmt.Println("They are not equal.")
	}
}
