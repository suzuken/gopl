package popcount

// pc[i] is the population count of i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

// PopCountLoop is loop version of PopCount
func PopCountLoop(x uint64) int {
	var ret byte
	for i := 0; i < 8; i++ {
		ret = ret + pc[byte(x>>uint64(i*8))]
	}
	return int(ret)
}

// PopCountAnd is bitwise AND version
func PopCountAnd(x uint64) int {
	return int(pc[byte(x&(x-1))])
}

// -> % go test -v -bench=. -benchmem github.com/suzuken/gopl/ch2/popcount64
// === RUN   TestPopCount
// --- PASS: TestPopCount (0.00s)
// PASS
// BenchmarkPopCount-8     300000000                4.17 ns/op            0 B/op          0 allocs/op
// BenchmarkPopCountLoop-8 100000000               10.7 ns/op             0 B/op          0 allocs/op
// ok      github.com/suzuken/gopl/ch2/popcount64  2.761s
