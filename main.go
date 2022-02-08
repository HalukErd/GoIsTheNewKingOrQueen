package main

import (
	"fmt"
	"strconv"
	"strings"
)

func mathString(str string) string {
	str = editBefore(str)
	str = noBracket(str)
	return str
}

func editBefore(str string) string {
	var strB strings.Builder
	strB.WriteString("(")
	for i, ch := range str {
		if ch == ')' && i != len(str)-1 && !strings.ContainsAny(str[i+1:i+2], "+-*/") {
			strB.WriteString(")*")
		} else if ch == '-' {
			strB.WriteString("+-")
		} else {
			strB.WriteRune(ch)
		}
	}
	strB.WriteString(")")
	return strB.String()
}

func noBracket(str string) string {
	var strB strings.Builder

	for i := 0; i < len(str); i++ {

		if str[i] == '(' {
			str = str[:i] + noBracket(str[i+1:])
			i--
		} else if str[i] == ')' {
			strB.Reset()
			strB.WriteString(calc(str[:i]))
			strB.WriteString(str[i+1:])
			return strB.String()
		} else {
			strB.WriteRune(rune(str[i]))
		}
	}
	return strB.String()
}

func calc(str string) string {
	if str[0] == '+' || str[0] == '-' {
		return str
	}
	sl := strings.Split(str, "+")
	var sum int
	for i := 0; i < len(sl); i++ {
		if strings.ContainsAny(sl[i], "*/") {
			sum += operate(sl[i])
		} else {
			atoi, _ := strconv.Atoi(sl[i])
			sum += atoi
		}
	}
	return strconv.Itoa(sum)
}
func operate(str string) int {
	result := 1
	lastIndex := 0
	var op rune
	isFirst := true
	for i, ch := range str {
		if isFirst && (ch == '*' || ch == '/') {
			op = ch
			lastIndex = i
			atoi, _ := strconv.Atoi(str[:i])
			result *= atoi
			isFirst = false
		} else if ch == '*' || ch == '/' {
			atoi, _ := strconv.Atoi(str[lastIndex+1 : i])
			result = miniOp(result, op, atoi)
			op = ch
			lastIndex = i
		} else if i == len(str)-1 {
			atoi, _ := strconv.Atoi(str[lastIndex+1 : i+1])
			result = miniOp(result, op, atoi)
		}
	}
	return result
}

func miniOp(a int, op rune, b int) int {
	if op == '*' {
		return a * b
	} else {
		return a / b
	}
}

func main() {
	fmt.Println(mathString("1+22*(3+1)(35-(33*2)+33)-2*(3)"))
}
