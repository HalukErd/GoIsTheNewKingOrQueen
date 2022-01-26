package main

import (
	"fmt"
	"regexp"
	"strings"
)

func LongestWord(sen string) string {
	var l int
	var lw string
	sl := strings.Split(sen, " ")

	for i := 0; i < len(sl); i++ {

		if lenCh(sl[i]) > l {
			lw = sl[i]
			l = lenCh(sl[i])
		}
	}
	// code goes here
	return lw

}

func lenCh(str string) int {
	var l int
	sl := strings.Split(str, "")

	for i := 0; i < len(sl); i++ {
		match, _ := regexp.MatchString("[a-zA-Z0-9]", sl[i])
		if match {
			l++
		}
	}
	return l
}

func main() {

	// do not modify below here, readline is our function
	// that properly reads in the input for you
	fmt.Println(LongestWord("Hello world123 567"))

}
