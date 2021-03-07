package stringprefix

import (
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {
	got := longestCommonPrefix([]string{"flow", "flower", "flight"})
	want := "fl"
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
