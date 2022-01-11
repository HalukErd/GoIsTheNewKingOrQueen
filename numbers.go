package main

type numbers []int

func newNumbers(size int) numbers {
	var s []int
	for i := 0; i < size; i++ {
		s = append(s, i)
	}
	return s
}

func (n numbers) considerAndPrintOddOrOven() {
	for _, number := range n {

		println(number, "is", isEvenOrOdd(number))
	}
}

func isEvenOrOdd(number int) string {
	if number%2 == 0 {
		return "even"
	} else {
		return "odd"
	}
}
