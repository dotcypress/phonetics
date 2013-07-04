// Copyright 2013 Vitaly Domnikov. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package phonetics

import "testing"

func TestMetaphoneEmptyString(t *testing.T) {
	if EncodeMetaphone("") != "" {
		t.Errorf("Encode with empty string should return empty string")
	}
}

func TestMetaphoneEncode(t *testing.T) {
	assertMetaphoneEquals(t, "Donald", "TNLT")
	assertMetaphoneEquals(t, "Zach", "SX")
	assertMetaphoneEquals(t, "Campbel", "KMPBL")
	assertMetaphoneEquals(t, "Cammmppppbbbeeelll", "KMPBL")
	assertMetaphoneEquals(t, "David", "TFT")
	assertMetaphoneEquals(t, "Wat", "WT")
	assertMetaphoneEquals(t, "What", "WT")
	assertMetaphoneEquals(t, "Gaspar", "KSPR")
	assertMetaphoneEquals(t, "ggaspar", "KSPR")
}

func assertMetaphoneEquals(t *testing.T, source string, target string) {
	if EncodeMetaphone(source) != target {
		t.Errorf("source doesn't match target. Input: %s, Result: %s, Target: %s", source, EncodeMetaphone(source), target)
	}
}
