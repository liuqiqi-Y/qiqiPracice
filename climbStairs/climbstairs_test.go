package climbStairs

import "testing"

func BenchmarkClimb(b *testing.B) {
	b.Run("动态规划方法", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			climbStairs1(10)
		}
	})
	b.Run("递归方法", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			climbStairs2(10)
		}
	})
}
func TestClimb(t *testing.T) {
	got := climbStairs1(3)
	want := 3
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
