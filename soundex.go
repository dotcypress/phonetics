// Copyright 2013 Vitaly Domnikov. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package phonetics

import (
	"strings"
)

// EncodeSoundex is a function to encode string with Soundex algorithm.
// Soundex is a phonetic algorithm for indexing names by sound, as pronounced in English.
func EncodeSoundex(word string) string {
	if word == "" {
		return "0000"
	}
	input := strings.ToLower(word)
	result := strings.ToUpper(input[0:1])
	code := ""
	lastCode := ""
	for _, rune := range input[1:] {
		switch rune {
		case 'b', 'f', 'p', 'v':
			code = "1"
		case 'c', 'g', 'j', 'k', 'q', 's', 'x', 'z':
			code = "2"
		case 'd', 't':
			code = "3"
		case 'l':
			code = "4"
		case 'm', 'n':
			code = "5"
		case 'r':
			code = "6"
		}
		if lastCode != code {
			lastCode = code
			result = result + lastCode
			if len(result) == 4 {
				break
			}
		}
	}
	return result + strings.Repeat("0", 4-len(result))
}

// DifferenceSoundex is a function to calculate difference between two strings with Soundex algorithm.
// Function returns a ranking on how similar two words are in percents.
func DifferenceSoundex(word1, word2 string) int {
	sum := differenceSoundex(word1, word2) + differenceSoundex(word2, word1)
	if sum == 0 {
		return 0
	}
	return sum / 2
}

func differenceSoundex(word1, word2 string) int {
	soundex1 := EncodeSoundex(word1)
	soundex2 := EncodeSoundex(word2)
	if soundex1 == soundex2 {
		return 100
	}
	result := 0
	if strings.Index(soundex2, soundex1[1:]) > -1 {
		result = 3
	} else if strings.Index(soundex2, soundex1[2:]) > -1 || strings.Index(soundex2, soundex1[1:3]) > -1 {
		result = 2
	} else {
		if strings.Index(soundex2, soundex1[1:2]) > -1 {
			result = result + 1
		}
		if strings.Index(soundex2, soundex1[2:3]) > -1 {
			result = result + 1
		}
		if strings.Index(soundex2, soundex1[3:4]) > -1 {
			result = result + 1
		}
	}
	if soundex1[0:1] == soundex2[0:1] {
		result = result + 1
	}
	return result * 25
}
