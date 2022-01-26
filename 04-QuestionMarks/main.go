package main

import (
	"fmt"
	"strconv"
	"strings"
)

func QuestionsMarks(str string) string {
	result := "false"
	success := true
	sl := strings.Split(str, "")
	digits := make(map[int]int)

	for i, ch := range sl {
		n, err := strconv.Atoi(ch)
		if err == nil {
			digits[i] = n
		}
	}
	for k, v := range digits {
		for kx, vx := range digits {
			if k < kx && v+vx == 10 {
				if checkThreeMuskeeters(k, kx, str) {
					result = "true"
				} else {
					success = false
					break
				}
			}
		}
	}
	if !success {
		result = "false"
	}
	// code goes here
	return result

}

func checkThreeMuskeeters(start int, end int, str string) bool {
	s := str[start+1 : end]
	count := 0
	for _, ch := range s {
		if ch == '?' {
			count++
		}
	}

	return count == 3
}

func main() {

	// do not modify below here, readline is our function
	// that properly reads in the input for you
	fmt.Println(QuestionsMarks("arrb6???4xxbl5???eee5"))

}
