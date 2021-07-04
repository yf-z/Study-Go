// exercise 2.3
package chap2Excer

import (
	"fmt"
	"time"
)

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))]+
	pc[byte(x>>(1*8))]+
	pc[byte(x>>(2*8))]+
	pc[byte(x>>(3*8))]+
	pc[byte(x>>(4*8))]+
	pc[byte(x>>(5*8))]+
	pc[byte(x>>(6*8))]+
	pc[byte(x>>(7*8))])
}

func PopCountCompare(x uint64) int {
	res := 0
	for i := 0; i < 8; i++ {
		res += int(pc[byte(x>>(i*8))])
	}
	return res
}

func Compare(x uint64) {
	start := time.Now().UnixNano()
	res1 := PopCount(x)
	end := time.Now().UnixNano()
	dur1 := float64(end-start)

	start = time.Now().UnixNano()
	res2 := PopCountCompare(x)
	end = time.Now().UnixNano()
	dur2 := float64(end-start)

	fmt.Printf("res1 = %d, duration1 = %g nano; res2 = %d, duration2 = %g nano\n", res1, dur1, res2, dur2)
}