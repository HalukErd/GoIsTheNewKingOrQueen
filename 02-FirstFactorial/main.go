package main

import "fmt"

func FirstFactorial(num int) int {
	if num == 1 {
		return 1
	}
	return num * FirstFactorial(num-1)
}

func main() {

	// do not modify below here, readline is our function
	// that properly reads in the input for you
	fmt.Println(FirstFactorial(8))

}
