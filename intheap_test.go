package heap

import (
	"fmt"
	"testing"
)

func TestHeap(t *testing.T) {
	h := []int{5, 7, 1, 3}
	Init(&h)
	Push(&h, 3)
	if len(h) != 5 {
		t.Errorf("push failed")
	}
	verify(t, &h, 0)
	min := Pop(&h)
	if min != 0 {
		t.Errorf("push failed")
	}
	fmt.Printf("minimum: %d\n", min)
	for len(h) > 0 {
		fmt.Printf("%d ", Pop(&h))
	}
	// Output:
	// minimum: 1
	// 1 2 3 5
}
