package main

import (
	"io"
	"os"
	"regexp"
)

func main() {
	for _, fn := range os.Args {
		matched, errMatch := regexp.MatchString("/*.txt", fn)
		if errMatch == nil && matched {
			f, err := os.Open(fn)
			if err == nil {
				io.Copy(os.Stdout, f)
			}
		}
	}
}
