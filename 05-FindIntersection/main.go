package main

import (
	"fmt"
	"strings"
)

func FindIntersection(strArr []string) string {

	sl1 := strings.Split(strArr[0], ", ")
	sl2 := strings.Split(strArr[1], ", ")

	m1 := make(map[string]bool)
	for _, v := range sl1 {
		m1[v] = true
	}

	sl := make([]string, 0)
	for _, v := range sl2 {
		if m1[v] {
			sl = append(sl, v)
		}
	}
	if len(sl) > 0 {
		return strings.Join(sl, ",")
	}
	return "false"
}

func main() {

	// do not modify below here, readline is our function
	// that properly reads in the input for you
	fmt.Println(FindIntersection([]string{"1, 3, 4, 7, 13", "1, 2, 4, 13, 15"}))

}
