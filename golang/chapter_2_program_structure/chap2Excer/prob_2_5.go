// exercise 2.5
package chap2Excer

import (
	"fmt"
	"time"
)

func PopCount5(x uint64) int {
	res := 0
	for ; x > 0; x, res = x&(x-1), res+1 {}
	return res
}

func Compare5(x uint64) {
	start := time.Now().UnixNano()
	res1 := PopCount(x)
	end := time.Now().UnixNano()
	dur1 := float64(end-start)

	start = time.Now().UnixNano()
	res2 := PopCount5(x)
	end = time.Now().UnixNano()
	dur2 := float64(end-start)

	fmt.Printf("res1 = %d, duration1 = %g nano; res2 = %d, duration2 = %g nano\n", res1, dur1, res2, dur2)
}