package main

import (
	"fmt"
	"strconv"
	"strings"
)

type root struct {
	self   int
	lchild int
	rchild int
}

func TreeConstructor(strArr []string) string {
	flag := true
	m := make(map[int]root)

	for _, str := range strArr {
		str = strings.Trim(str, "(")
		str = strings.Trim(str, ")")
		sl := strings.Split(str, ",")
		c, _ := strconv.Atoi(sl[0])
		p, _ := strconv.Atoi(sl[1])
		r := m[p]
		if c < p {
			if r.lchild != 0 {
				flag = false
				break
			}
			r.lchild = c
		} else if c > p {
			if r.rchild != 0 {
				flag = false
				break
			}
			r.rchild = c
		} else {
			flag = false
			break
		}
		r.self = p
		m[p] = r
	}
	if flag && mainRootCheck(m) {
		return "true"
	}
	return "false"

}

func mainRootCheck(m map[int]root) bool {
	mainRoot := 0
	rootCheck := make(map[int]bool)
	for _, v := range m {
		rootCheck[v.rchild] = true
		rootCheck[v.lchild] = true
	}
	for k, _ := range m {
		if !rootCheck[k] {
			mainRoot++
		}
	}
	return mainRoot == 1
}

func main() {

	// do not modify below here, readline is our function
	// that properly reads in the input for you
	fmt.Println(TreeConstructor([]string{"(1,2)", "(2,4)", "(7,2)"}))

}
