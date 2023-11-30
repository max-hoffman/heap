package heap

import (
	goheap "container/heap"
	"errors"
	"fmt"
)

var ErrUnsupportedHeap = errors.New("dynamic heap support not found for type:")

func Init(h interface{}) error {
	// heapify
	switch h := h.(type) {
	case goheap.Interface:
		goheap.Init(h)
	default:
		if !IsSupported(h) {
			return fmt.Errorf("%w: %T", ErrUnsupportedHeap, h)
		}
		n := Len(h)
		for i := n/2 - 1; i >= 0; i-- {
			down(h, i, n)
		}
	}
	return nil
}

// Push pushes the element x onto the heap.
// The complexity is O(log n) where n = h.Len().
func Push(h interface{}, x any) interface{} {
	switch h := h.(type) {
	case goheap.Interface:
		goheap.Push(h, x)
	default:
		Append(h, x)
		up(h, Len(h)-1)
	}
	return nil
}

func Append(h interface{}, x any) {
	switch h := h.(type) {
	case *[]int:
		*h = append(*h, x.(int))
	case *[]string:
		*h = append(*h, x.(string))
	default:
		panic(fmt.Sprintf("unsupported heap type: %T", h))
	}
}

// Pop removes and returns the minimum element (according to Less) from the heap.
// The complexity is O(log n) where n = h.Len().
// Pop is equivalent to Remove(h, 0).
func Pop(h interface{}) any {
	switch h := h.(type) {
	case goheap.Interface:
		return goheap.Pop(h)
	default:
		n := Len(h) - 1
		Swap(h, 0, n)
		down(h, 0, n)
		return PopLast(h)
	}
}

func PopLast(h interface{}) any {
	switch h := h.(type) {
	case *[]int:
		arr := *h
		ret := arr[len(arr)-1]
		*h = arr[:len(arr)-1]
		return ret
	case *[]string:
		arr := *h
		ret := arr[len(arr)-1]
		*h = arr[:len(arr)-1]
		return ret
	default:
		panic(fmt.Sprintf("unsupported type: %T", h))
	}
}

func Len(h interface{}) int {
	switch h := h.(type) {
	case *[]int:
		return len(*h)
	case *[]string:
		return len(*h)
	default:
		panic(fmt.Sprintf("unsupported type: %T", h))
	}
}

// Remove removes and returns the element at index i from the heap.
// The complexity is O(log n) where n = h.Len().
func Remove(h interface{}, i int) any {
	switch h := h.(type) {
	case goheap.Interface:
		return goheap.Remove(h, i)
	default:
		n := Len(h) - 1
		if n != i {
			Swap(h, i, n)
			if !down(h, i, n) {
				up(h, i)
			}
		}
		return PopLast(h)
	}
}

// Fix re-establishes the heap ordering after the element at index i has changed its value.
// Changing the value of the element at index i and then calling Fix is equivalent to,
// but less expensive than, calling Remove(h, i) followed by a Push of the new value.
// The complexity is O(log n) where n = h.Len().
func Fix(h interface{}, i int) {
	switch h := h.(type) {
	case goheap.Interface:
		goheap.Fix(h, i)
	default:
		if !down(h, i, Len(h)) {
			up(h, i)
		}
	}
}

func Swap(h interface{}, i, j int) {
	switch h := h.(type) {
	case *[]int:
		arr := *h
		arr[i], arr[j] = arr[j], arr[i]
	case *[]string:
		arr := *h
		arr[i], arr[j] = arr[j], arr[i]
	default:
		panic(fmt.Sprintf("unsupported type: %T", h))
	}
}

func up(h interface{}, j int) {
	for {
		i := (j - 1) / 2 // parent
		if i == j || !Less(h, j, i) {
			break
		}
		Swap(h, i, j)
		j = i
	}
}

func down(h interface{}, i0, n int) bool {
	i := i0
	for {
		j1 := 2*i + 1
		if j1 >= n || j1 < 0 { // j1 < 0 after int overflow
			break
		}
		j := j1 // left child
		if j2 := j1 + 1; j2 < n && Less(h, j2, j1) {
			j = j2 // = 2*i + 2  // right child
		}
		if !Less(h, j, i) {
			break
		}
		Swap(h, i, j)
		i = j
	}
	return i > i0
}

func IsSupported(h interface{}) bool {
	switch h.(type) {
	case *[]int:
		return true
	default:
		return false
	}
}

func Less(h interface{}, i, j int) bool {
	switch h := h.(type) {
	case goheap.Interface:
		return h.Less(i, j)
	case *[]int:
		arr := *h
		return arr[i] < arr[j]
	case *[]string:
		arr := *h
		return arr[i] < arr[j]
	default:
		panic(fmt.Sprintf("unsupported heap element type: %T", h))
	}
}
