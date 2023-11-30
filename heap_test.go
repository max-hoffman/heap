// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package heap

import (
	"math/rand"
	"testing"
)

type myHeap []int

func (h *myHeap) Less(i, j int) bool {
	return (*h)[i] < (*h)[j]
}

func (h *myHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *myHeap) Len() int {
	return len(*h)
}

func (h *myHeap) Pop() (v any) {
	*h, v = (*h)[:len(*h)-1], (*h)[len(*h)-1]
	return
}

func (h *myHeap) Push(v any) {
	*h = append(*h, v.(int))
}

func verify(t *testing.T, h *[]int, i int) {
	t.Helper()
	n := len(*h)
	j1 := 2*i + 1
	j2 := 2*i + 2
	if j1 < n {
		if Less(h, j1, i) {
			arr := *h
			t.Errorf("heap invariant invalidated [%d] = %d > [%d] = %d", i, arr[i], j1, arr[j1])
			return
		}
		verify(t, h, j1)
	}
	if j2 < n {
		if Less(h, j2, i) {
			arr := *h
			t.Errorf("heap invariant invalidated [%d] = %d > [%d] = %d", i, arr[i], j1, arr[j2])
			return
		}
		verify(t, h, j2)
	}
}

func TestInit0(t *testing.T) {
	h := new([]int)
	for i := 20; i > 0; i-- {
		Push(h, 0) // all elements are the same
	}
	Init(h)
	verify(t, h, 0)

	for i := 1; len(*h) > 0; i++ {
		x := Pop(h).(int)
		verify(t, h, 0)
		if x != 0 {
			t.Errorf("%d.th pop got %d; want %d", i, x, 0)
		}
	}
}

func TestInit1(t *testing.T) {
	h := new([]int)
	for i := 20; i > 0; i-- {
		Push(h, i) // all elements are different
	}
	Init(h)
	verify(t, h, 0)

	for i := 1; len(*h) > 0; i++ {
		x := Pop(h).(int)
		verify(t, h, 0)
		if x != i {
			t.Errorf("%d.th pop got %d; want %d", i, x, i)
		}
	}
}

func Test(t *testing.T) {
	h := new([]int)
	verify(t, h, 0)

	for i := 20; i > 10; i-- {
		Push(h, i)
	}
	Init(h)
	verify(t, h, 0)

	for i := 10; i > 0; i-- {
		Push(h, i)
		verify(t, h, 0)
	}

	for i := 1; len(*h) > 0; i++ {
		x := Pop(h).(int)
		if i < 20 {
			Push(h, 20+i)
		}
		verify(t, h, 0)
		if x != i {
			t.Errorf("%d.th pop got %d; want %d", i, x, i)
		}
	}
}

func TestRemove0(t *testing.T) {
	h := new([]int)
	for i := 0; i < 10; i++ {
		Push(h, i)
	}
	verify(t, h, 0)

	for Len(h) > 0 {
		i := Len(h) - 1
		x := Remove(h, i).(int)
		if x != i {
			t.Errorf("Remove(%d) got %d; want %d", i, x, i)
		}
		verify(t, h, 0)
	}
}

func TestRemove1(t *testing.T) {
	h := new([]int)
	for i := 0; i < 10; i++ {
		Push(h, i)
	}
	verify(t, h, 0)

	for i := 0; len(*h) > 0; i++ {
		x := Remove(h, 0).(int)
		if x != i {
			t.Errorf("Remove(0) got %d; want %d", x, i)
		}
		verify(t, h, 0)
	}
}

func TestRemove2(t *testing.T) {
	N := 10

	h := new([]int)
	for i := 0; i < N; i++ {
		Push(h, i)
	}
	verify(t, h, 0)

	m := make(map[int]bool)
	for len(*h) > 0 {
		m[Remove(h, (len(*h)-1)/2).(int)] = true
		verify(t, h, 0)
	}

	if len(m) != N {
		t.Errorf("len(m) = %d; want %d", len(m), N)
	}
	for i := 0; i < len(m); i++ {
		if !m[i] {
			t.Errorf("m[%d] doesn't exist", i)
		}
	}
}

func BenchmarkDup(b *testing.B) {
	const n = 10000
	h := make([]int, 0, n)
	for i := 0; i < b.N; i++ {
		for j := 0; j < n; j++ {
			Push(&h, 0) // all elements are the same
		}
		for len(h) > 0 {
			Pop(&h)
		}
	}
}

func TestFix(t *testing.T) {
	h := new([]int)
	verify(t, h, 0)

	for i := 200; i > 0; i -= 10 {
		Push(h, i)
	}
	verify(t, h, 0)

	if (*h)[0] != 10 {
		t.Fatalf("Expected head to be 10, was %d", (*h)[0])
	}
	(*h)[0] = 210
	Fix(h, 0)
	verify(t, h, 0)

	for i := 100; i > 0; i-- {
		elem := rand.Intn(Len(h))
		if i&1 == 0 {
			(*h)[elem] *= 2
		} else {
			(*h)[elem] /= 2
		}
		Fix(h, elem)
		verify(t, h, 0)
	}
}
