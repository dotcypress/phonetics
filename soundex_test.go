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
		t.Errorf("source doesn't match target. %s -> %s", source, EncodeSoundex(source))
	}
}
