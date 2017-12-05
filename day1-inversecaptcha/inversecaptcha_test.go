package inversecaptcha

import (
	"testing"
)

func TestInverseCaptcha(t *testing.T) {
	for _, tc := range addCases {
		in := tc.in
		want := tc.want
		got := InverseCaptcha(in)

		if got != want {
			t.Errorf(`FAIL: %s
InverseCaptcha(%s)
= %d
want %d`, tc.description, tc.in, got, tc.want)
		}
		t.Log("PASS:", tc.description)
	}
	t.Log("Tested", len(addCases), "cases.")
}
