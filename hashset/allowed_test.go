package hashset

import "testing"

func TestHashset(t *testing.T) {
	test := New()
	test.Add("testing1")

	if !test.Contains("testing1") || test.Contains("testing2") {
		t.Fatal()
	}
}
