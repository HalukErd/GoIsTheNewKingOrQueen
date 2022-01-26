package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func CodelandUsernameValidation(str string) string {

	match, _ := regexp.MatchString("[a-zA-Z]+[a-zA-Z0-9_]*[a-zA-Z0-9]$", str)

	return strconv.FormatBool(match)

}

func main() {

	// do not modify below here, readline is our function
	// that properly reads in the input for you
	fmt.Println(CodelandUsernameValidation("aaaaaaaaaa"))

}
