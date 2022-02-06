package main

import (
	"fmt"
	"math/big"
)

func extraLongFactorials(n int32) {

	theBig := big.NewInt(int64(1))
	for i := 2; int32(i) <= n; i++ {
		theBig.Mul(theBig, big.NewInt(int64(i)))
	}
	fmt.Printf("%v", theBig)

	nines := make([]int64, 1, 100)
	nines[0] = 1
	const bil int64 = 1000000000

	for i := 1; i <= int(n); i++ {

		lenNines := len(nines)
		var left int64
		for j := 0; j < lenNines; j++ {
			nines[j] = nines[j]*int64(i) + left
			if nines[j] > bil {
				mod := nines[j] % bil
				left = nines[j] / bil
				nines[j] = mod
				if j == lenNines-1 {
					nines = append(nines, left)
					left = 0
				}
			} else {
				left = 0
			}
		}
	}
	fmt.Printf("%d", nines[len(nines)-1])
	for i := len(nines) - 2; i >= 0; i-- {
		fmt.Printf("%09d", nines[i])
	}
}

func main() {
	extraLongFactorials(10)
}
