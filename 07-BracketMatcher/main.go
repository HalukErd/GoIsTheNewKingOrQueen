package main

import (
	"fmt"
)

func BracketMatcher(str string) string {
	result := 0
	for _, ch := range str {
		if ch == '(' {
			result++
		} else if ch == ')' {
			result--
		}
		if result < 0 {
			return "0"
		}
	}
	if result == 0 {
		return "1"
	}
	return "0"
}

func main() {

	// do not modify below here, readline is our function
	// that properly reads in the input for you
	fmt.Println(BracketMatcher("hello()"))

}
