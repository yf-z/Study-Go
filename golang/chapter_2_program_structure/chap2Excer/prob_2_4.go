// exercise 2.4
package chap2Excer

import (
	"fmt"
	"time"
)

func PopCountDigits(x uint64) int {
	res := 0
	for ; x > 0; x = x >> 1 {
		res += int(x & 1)
	}
	return res
}

func Compare3(x uint64) {
	start := time.Now().UnixNano()
	res1 := PopCount(x)
	end := time.Now().UnixNano()
	dur1 := float64(end-start)

	start = time.Now().UnixNano()
	res2 := PopCountDigits(x)
	end = time.Now().UnixNano()
	dur2 := float64(end-start)

	fmt.Printf("res1 = %d, duration1 = %g nano; res2 = %d, duration2 = %g nano\n", res1, dur1, res2, dur2)
}