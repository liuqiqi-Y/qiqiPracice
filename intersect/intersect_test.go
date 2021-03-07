package intersect

import (
	"reflect"
	"testing"
)

func TestIntersect(t *testing.T) {
	got := intersect([]int{4, 9, 5}, []int{9, 4, 9, 8, 4})
	want := []int{9, 4}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
