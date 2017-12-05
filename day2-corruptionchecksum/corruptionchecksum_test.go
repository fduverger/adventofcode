package corruptionchecksum

import (
	"testing"
)

func TestCorruptionChecksum(t *testing.T) {
	for _, tc := range addCases {
		in := tc.in
		want := tc.want
		got := CorruptionChecksum(in)

		if got != want {
			t.Errorf(`FAIL: %s
CorruptionChecksum(%v)
= %d
want %d`, tc.description, tc.in, got, tc.want)
		}
		t.Log("PASS:", tc.description)
	}
	t.Log("Tested", len(addCases), "cases.")
}
