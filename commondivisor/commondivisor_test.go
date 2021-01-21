package commondivisor

import "testing"

func TestCommonDivisor(t *testing.T) {
	got := commonDivisor2(9, 6)
	want := 3
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func BenchmarkCommonDivisor(b *testing.B) {
	b.Run("非递归方法", func(b *testing.B) {
		for i := 1; i < b.N; i++ {
			commonDivisor1(9, 6)
		}
	})
	b.Run("递归方法", func(b *testing.B) {
		for i := 1; i < b.N; i++ {
			commonDivisor2(9, 6)
		}
	})
}
