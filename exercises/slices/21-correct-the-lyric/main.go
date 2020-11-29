package main

import (
	"fmt"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Correct the lyric
//
//  You have a slice that contains the words of Beatles'
//  legendary song: Yesterday. However, the order of the
//  words are incorrect.
//
// CURRENT OUTPUT
//  [all my troubles seemed so far away oh i believe in yesterday now it looks as though they are here to stay]
//
// EXPECTED OUTPUT
//  [yesterday all my troubles seemed so far away now it looks as though they are here to stay oh i believe in yesterday]
//
// STEPS
//  INITIAL SLICE:
//    [all my troubles seemed so far away oh i believe in yesterday now it looks as though they are here to stay]
//
//  1. Prepend "yesterday" to the `lyric` slice.
//     RESULT SHOULD BE:
//     [yesterday all my troubles seemed so far away oh i believe in yesterday now it looks as though they are here to stay]
//
//  2. Put the words to the correct positions in the `lyric` slice.
//     RESULT SHOULD BE:
//     [yesterday all my troubles seemed so far away now it looks as though they are here to stay oh i believe in yesterday]
//
//  3. Print the `lyric` slice.
// ---------------------------------------------------------

func main() {
	lyric := strings.Fields(`all my troubles seemed so far away oh i believe in yesterday now it looks as though they are here to stay`)
	lyric = append([]string{"yesterday"}, lyric...)
	swapLyrics := append(lyric[13:], lyric[8:13]...)
	lyric = append(lyric[:8], swapLyrics...)
	fmt.Printf("%s\n", lyric)
}
