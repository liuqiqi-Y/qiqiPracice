package listreverse

import "fmt"

type node struct {
	val  int
	next *node
}

func makeList(n int) *node {
	head := node{}
	p := &head
	for i := 1; i <= n; i++ {
		v := node{
			val: i,
		}
		p.next = &v
		p = &v
	}
	return &head
}
func reverse(head *node) *node {
	if head.next == nil || head == nil {
		return head
	}
	h := reverse(head.next)
	head.next.next = head
	head.next = nil
	return h
}
func printList(list *node) {
	p := list.next
	for p != nil {
		fmt.Printf("%d --> ", p.val)
		p = p.next
	}
}
