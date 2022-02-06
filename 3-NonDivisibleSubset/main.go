package main

import (
	"fmt"
)

func nonDivisibleSubset(k int32, s []int32) int32 {
	m := make(map[int32]int32)
	for _, i := range s {
		m[i%k]++
	}
	var result int32
	for i := int32(1); i <= k/2; i++ {

		if m[i] > m[k-i] {
			result += m[i]
		} else if m[k-i] > m[i] {
			result += m[k-i]
		} else if m[i] > 0 && i == k-i {
			result++
		}
		
	}
	if m[0] > 0 {
		result++
	}
	return result
}
func main() {
	subset := nonDivisibleSubset(7, []int32{278, 576, 496, 727, 410, 124, 338, 149, 209, 702, 282, 718, 771, 575, 436})
	fmt.Println(subset)
	subset = nonDivisibleSubset(9, []int32{422346306, 940894801, 696810740, 862741861, 85835055, 313720373})
	fmt.Println(subset)
	subset = nonDivisibleSubset(4, []int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	fmt.Println(subset)
}
