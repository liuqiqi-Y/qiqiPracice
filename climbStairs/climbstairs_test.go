package climbStairs

import "testing"

func TestClimb(t *testing.T) {
	got := climbStairs1(3)
	want := 3
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
