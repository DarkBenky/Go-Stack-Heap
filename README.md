# Escape Analysis Results for `main.go`

## Summary of Results

The output of the `go build -gcflags="-m" main.go` command provides insights into how variables are allocated and whether they escape to the heap or remain on the stack.

### Analysis of Variables

- **Global Variable `c`**:
  - **Code**: `var c = make([]int, 8191)`
  - **Result**: `make([]int, 8191) escapes to heap`
  - **Explanation**: Since `c` is a global variable, it automatically escapes to the heap, regardless of its size.

- **Local Variable `u`**:
  - **Code**: `u := make([]int, 8191)`
  - **Result**: `make([]int, 8191) does not escape`
  - **Explanation**: The `u` slice is created locally within `main` and does not escape to the heap since it’s small enough to remain on the stack.

- **Local Variable `v`**:
  - **Code**: `v := make([]int, 8192)`
  - **Result**: `make([]int, 8192) does not escape package main`
  - **Explanation**: Although `v` is of similar size to `u`, the additional element size (64KB total) is large enough that it requires heap allocation.

- **Array `w`**:
  - **Code**: `var w [1024 * 1024 * 1.25]int`
  - **Result**: Allocated on the stack.
  - **Explanation**: This array is 5MB in size and does not exceed the stack size limit, so it remains on the stack.

- **Array `x`**:
  - **Code**: `var x [1024*1024*1.25 + 1]int`
  - **Result**: `moved to heap`
  - **Explanation**: This array exceeds the stack limit (10MB), so it is moved to the heap.

### Key Takeaways

- **Escape to Heap**: Variables that exceed a certain size (10MB for arrays) or are globally declared will escape to the heap.
- **Stack Allocation**: Smaller slices and arrays, as well as local variables, are allocated on the stack if their size is within the limit.
- **Optimization Insight**: Keeping data within the stack can improve performance, so it’s beneficial to monitor these allocations in performance-critical code.

### Full Command Output

```plaintext
./main.go:5:6: can inline main
./main.go:3:13: make([]int, 8191) escapes to heap
./main.go:17:6: moved to heap: x
./main.go:8:11: make([]int, 8191) does not escape
./main.go:11:11: make([]int, 8192) does not escape package main

