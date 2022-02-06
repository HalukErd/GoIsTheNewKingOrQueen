package main

import (
	"fmt"
	"strconv"
	"strings"
)

var m = make(map[string]bool)

func queensAttack(n int32, k int32, r_q int32, c_q int32, obstacles [][]int32) int32 {

	for i := int32(0); i < k; i++ {
		obsX := strconv.Itoa(int(obstacles[i][0]))
		obsY := strconv.Itoa(int(obstacles[i][1]))
		m[strings.Join([]string{obsX, obsY}, ",")] = true
	}
	flags := [8]bool{true, true, true, true, true, true, true, true}
	var result int32
	for i := int32(1); i < n; i++ {
		direction(&flags[0], r_q, c_q+i, &result, n, obstacles)
		direction(&flags[1], r_q, c_q-i, &result, n, obstacles)
		direction(&flags[2], r_q+i, c_q, &result, n, obstacles)
		direction(&flags[3], r_q-i, c_q, &result, n, obstacles)
		direction(&flags[4], r_q+i, c_q+i, &result, n, obstacles)
		direction(&flags[5], r_q-i, c_q-i, &result, n, obstacles)
		direction(&flags[6], r_q-i, c_q+i, &result, n, obstacles)
		direction(&flags[7], r_q+i, c_q-i, &result, n, obstacles)
	}
	return result
}

func direction(flag *bool, rqD int32, cqD int32, result *int32, n int32, obs [][]int32) {
	if !*flag {
		return
	}
	if rqD < 1 || rqD > n || cqD < 1 || cqD > n {
		*flag = false
		return
	}
	obsX := strconv.Itoa(int(rqD))
	obsY := strconv.Itoa(int(cqD))
	if m[strings.Join([]string{obsX, obsY}, ",")] {
		*flag = false
		return
	}
	*result = *result + 1
}

func main() {
	attack := queensAttack(5, 3, 4, 3, [][]int32{{5, 5}, {4, 2}, {2, 3}})
	fmt.Println(attack)
	attack2 := queensAttack(4, 0, 4, 4, [][]int32{})
	fmt.Println(attack2)
	//5 3
	//4 3
	//5 5
	//4 2
	//2 3
}
