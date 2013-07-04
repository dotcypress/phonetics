// Copyright 2013 Vitaly Domnikov. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package phonetics

import (
	"strings"
)

// EncodeMetaphone is a function to encode string with Metaphone algorithm.
// Metaphone is a phonetic algorithm, published by Lawrence Philips in 1990, for indexing words by their English pronunciation.
// With Michael Kuhn modification (mkuhn@rhlab.UUCP)
func EncodeMetaphone(word string) string {
	if word == "" {
		return ""
	}
	word = strings.ToUpper(word)
	word = removeDuplicates(word)
	wordLen := len(word)
	if wordLen > 1 {
		switch word[0:2] {
		case "PN", "AE", "KN", "GN", "WR":
			word = word[1:]
		case "WH":
			word = "W" + word[2:]
		}
		if word[0:1] == "X" {
			word = "W" + word[1:]
		}
	}

	result := ""
	for i, rune := range word {
		switch rune {
		case 'B':
			{
				if i != wordLen-1 || safeSubString(word, i-1, 2) != "MB" {
					result = result + "B"
				}
			}
		case 'C':
			{
				if safeSubString(word, i, 3) == "CIA" || safeSubString(word, i, 2) == "CH" {
					result = result + "X"
				} else if safeSubString(word, i, 2) == "CI" || safeSubString(word, i, 2) == "CE" || safeSubString(word, i, 2) == "CY" {
					result = result + "S"
				} else if safeSubString(word, i-1, 3) != "SCI" || safeSubString(word, i-1, 3) != "SCE" || safeSubString(word, i-1, 3) != "SCY" {
					result = result + "K"
				}
			}
		case 'D':
			{
				if safeSubString(word, i, 3) == "DGE" || safeSubString(word, i, 3) == "DGY" || safeSubString(word, i, 3) == "DGI" {
					result = result + "J"
				} else {
					result = result + "T"
				}
			}
		case 'F':
			result = result + "F"
		case 'G':
			{
				prev := safeSubString(word, i+1, 1)
				if (safeSubString(word, i, 2) == "GH" && !isVowel(safeSubString(word, i+2, 1))) ||
					safeSubString(word, i, 2) == "GN" ||
					safeSubString(word, i, 4) == "GNED" ||
					safeSubString(word, i, 3) == "GDE" ||
					safeSubString(word, i, 3) == "GDY" ||
					safeSubString(word, i, 3) == "GDI" {
				} else if prev == "I" || prev == "E" || prev == "Y" {
					result = result + "J"
				} else {
					result = result + "K"
				}
			}
		case 'H':
			{
				if !isVowel(safeSubString(word, i+1, 1)) &&
					safeSubString(word, i-2, 2) != "CH" &&
					safeSubString(word, i-2, 2) != "SH" &&
					safeSubString(word, i-2, 2) != "PH" &&
					safeSubString(word, i-2, 2) != "TH" &&
					safeSubString(word, i-2, 2) != "GH" {
					result = result + "H"
				}
			}
		case 'J':
			result = result + "J"
		case 'K':
			{
				if safeSubString(word, i-1, 1) != "C" {
					result = result + "K"
				}
			}
		case 'L':
			result = result + "L"
		case 'M':
			result = result + "M"
		case 'N':
			result = result + "N"
		case 'P':
			{
				if safeSubString(word, i+1, 1) == "H" {
					result = result + "F"
				} else {
					result = result + "P"
				}
			}
		case 'Q':
			result = result + "K"
		case 'R':
			result = result + "R"
		case 'S':
			{
				if safeSubString(word, i+1, 1) == "H" || safeSubString(word, i, 3) == "SIO" || safeSubString(word, i, 3) == "SIA" {
					result = result + "X"
				} else {
					result = result + "S"
				}
			}
		case 'T':
			{
				if safeSubString(word, i, 3) == "TIO" || safeSubString(word, i, 3) == "TIA" {
					result = result + "X"
				} else if safeSubString(word, i+1, 1) == "H" {
					result = result + "0"
				} else if safeSubString(word, i, 3) != "TCH" {
					result = result + "T"
				}
			}
		case 'V':
			result = result + "F"
		case 'W':
			{
				if isVowel(safeSubString(word, i+1, 1)) {
					result = result + "W"
				}
			}
		case 'X':
			result = result + "KS"
		case 'Y':
			{
				if isVowel(safeSubString(word, i+1, 1)) {
					result = result + "Y"
				}
			}
		case 'Z':
			result = result + "S"
		}
	}
	return result
}

func safeSubString(word string, start, count int) string {
	wordLen := len(word)
	if start < 0 {
		start = 0
		count = count + start
	}
	if start+count > wordLen {
		count = wordLen - start
	}
	return word[start : start+count]
}

func isVowel(char string) bool {
	return strings.Index("AEIOU", char) > -1
}

func removeDuplicates(word string) string {
	previousChar := []rune(word)[0]
	result := string(previousChar)
	for _, rune := range word[1:] {
		if rune != previousChar || rune == 'C' {
			result = result + string(rune)
		}
		previousChar = rune
	}
	return result
}
