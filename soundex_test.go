// Copyright 2013 Vitaly Domnikov. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package phonetics

import "testing"

func TestEmptyString(t *testing.T) {
	if EncodeSoundex("") != "0000" {
		t.Errorf("Encode with empty string should return 0000")
	}
}

func TestEncode(t *testing.T) {
	assertEquals(t, "Donald", "D543")
	assertEquals(t, "Zach", "Z200")
	assertEquals(t, "Campbel", "C514")
	assertEquals(t, "Cammmppppbbbeeelll", "C514")
	assertEquals(t, "David", "D130")
}

func assertEquals(t *testing.T, source string, target string) {
	if EncodeSoundex(source) != target {
		t.Errorf("source doesn't match target. Input: %s, Result: %s, Target: %s", source, EncodeSoundex(source), target)
	}
}

func TestDifference(t *testing.T) {
	assertDifference(t, "Zach", "Zac", 100)
	assertDifference(t, "Lake", "Bake", 75)
	assertDifference(t, "Brad", "Lad", 50)
	assertDifference(t, "Horrible", "Great", 25)
	assertDifference(t, "Mike", "Jeremy", 37)
}

func assertDifference(t *testing.T, word1 string, word2 string, rank int) {
	if DifferenceSoundex(word1, word2) != rank {
		t.Errorf("difference doesn't match target. Input: (%s, %s), Result: %d, Target: %d", word1, word2, DifferenceSoundex(word1, word2), rank)
	}
}
