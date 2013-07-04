// Copyright 2013 Vitaly Domnikov. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package phonetics

import "testing"

func TestSoundexEmptyString(t *testing.T) {
	if EncodeSoundex("") != "0000" {
		t.Errorf("Encode with empty string should return 0000")
	}
}

func TestSoundexEncode(t *testing.T) {
	assertSoundexEquals(t, "Donald", "D543")
	assertSoundexEquals(t, "Zach", "Z200")
	assertSoundexEquals(t, "Campbel", "C514")
	assertSoundexEquals(t, "Cammmppppbbbeeelll", "C514")
	assertSoundexEquals(t, "David", "D130")
}

func TestSoundexDifference(t *testing.T) {
	assertSoundexDifference(t, "Zach", "Zac", 100)
	assertSoundexDifference(t, "Lake", "Bake", 75)
	assertSoundexDifference(t, "Brad", "Lad", 50)
	assertSoundexDifference(t, "Horrible", "Great", 25)
	assertSoundexDifference(t, "Mike", "Jeremy", 37)
}

func assertSoundexDifference(t *testing.T, word1 string, word2 string, rank int) {
	if DifferenceSoundex(word1, word2) != rank {
		t.Errorf("difference doesn't match target. Input: (%s, %s), Result: %d, Target: %d", word1, word2, DifferenceSoundex(word1, word2), rank)
	}
}

func assertSoundexEquals(t *testing.T, source string, target string) {
	if EncodeSoundex(source) != target {
		t.Errorf("source doesn't match target. Input: %s, Result: %s, Target: %s", source, EncodeSoundex(source), target)
	}
}
