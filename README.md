# Golang Default Heaps

Toy repo with implicit support for default slice types. The main
drawback is having to use slice pointers.

Currently supported:
- `*[]int`
- `*[]string`

Helper method `IsSupported` will return `false` if this library cannot
heap sort the type. A bit of effort codegen'ing default slice
harnesses would make it easy to support an arbitrary list of types.

Implementation falls back to standard library for types that implement
`heap.Interface` (standard library heap, not this one). Heap helper
methods will panic if the type passed `!IsSupported` and does not
implement `heap.Interface`.

Usage:
```golang
	h := []int{5, 7, 1, 3}
	Init(&h)
	Push(&h, 3)
	min := Pop(&h)
	fmt.Printf("minimum: %d\n", Pop(&h))
	for len(h) > 0 {
		fmt.Printf("%d ", Pop(&h))
	}
	// Output:
	// minimum: 1
	// 1 2 3 5
```
