package main

// type staticArray struct {
// 	data [128]int
// }

// type dynamicStuckArray struct {
// 	max    int
// 	arrays []staticArray
// }

// func (d *dynamicStuckArray) set(index int, value int) {
// 	for index >= d.max {
// 		d.arrays = append(d.arrays, staticArray{})
// 		d.max = d.max + len(d.arrays[0].data)
// 	}
// 	d.arrays[index/len(d.arrays[0].data)].data[index%len(d.arrays[0].data)] = value
// }

// func (d dynamicStuckArray) index(i int) int {
// 	if i >= d.max {
// 		panic("index out of range")
// 	}
// 	return d.arrays[i/len(d.arrays[0].data)].data[i%len(d.arrays[0].data)]
// }

var c = make([]int, 8191)

func main() {
	// d := dynamicStuckArray{}

	// start := time.Now()

	// d.set(25, 1)

	// for i := 0; i < 32768; i++ {
	// 	for j := 0; j < 20; j++ {
	// 		// d.set(j)
	// 		_ = d.index(j)
	// 	}
	// }

	// elapsed := time.Since(start)
	// fmt.Println(elapsed, "implementing dynamic stuck array")

	// v := make([]int, 256_000)
	// start = time.Now()

	// for i := 0; i < 32768; i++ {
	// 	for j := 0; j < 20; j++ {
	// 		// v[j] = j
	// 		_ = v[j]
	// 	}
	// }

	// elapsed = time.Since(start)
	// fmt.Println(elapsed, "using slice")

	// z := staticArray{}
	// _ = z

	_ = c

	u := make([]int, 8191) // Does not escape to heap
	_ = u

	v := make([]int, 8192) // Escapes to heap = 64kb
	_ = v

	var w [1024 * 1024 * 1.25]int
	_ = w

	var x [1024*1024*1.25 + 1]int // moved to heap > 10 mb
	_ = x
}
