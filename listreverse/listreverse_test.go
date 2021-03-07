package listreverse

import (
	"testing"
)

func TestMakeList(t *testing.T) {
	list := makeList(5)

	p := reverse(list.next)
	for p != nil {
		t.Errorf("%d --> ", p.val)
		p = p.next
	}
	t.Error("success")
}
