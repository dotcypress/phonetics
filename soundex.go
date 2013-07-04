// Copyright 2013 Vitaly Domnikov. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package phonetics

import (
	"strings"
)

// EncodeSoundex is a function to encode string options with Soundex algorithm.
// Soundex is a phonetic algorithm for indexing names by sound, as pronounced in English
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
