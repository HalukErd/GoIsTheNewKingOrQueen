package main

import (
	"fmt"
	"strings"
)

func FirstReverse(str string) string {
	length := len(str)
	sl := strings.Split(str, "")
	for i := 0; i < length/2; i++ {
		sl[i], sl[length-i-1] = sl[length-i-1], sl[i]
	}

	return strings.Join(sl, "")

}

func main() {

	// do not modify below here, readline is our function
	// that properly reads in the input for you
	fmt.Println(FirstReverse("coderbyte"))

}
