package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	s := os.Args[1]
	file, err := os.Open("p.txt")

	url1 := `<a href="https://bitbucket.org/avansashybristeam/avansas/branches/?search=%v">  ______  Search:%v  ______  </a>`
	url2 := `<a href="https://bitbucket.org/avansashybristeam/avansas/branch/master?dest=%v#diff">  ______  Merge Master %v  ______  </a>`
	url3 := `<a href="https://bitbucket.org/avansashybristeam/avansas/branch/%v?dest=%v">  ______  PR %v to %v  ______  </a>`

	if err != nil {
		log.Panic(err)
	}

	scnnr := bufio.NewScanner(file)

	for scnnr.Scan() {
		branch := scnnr.Text()
		fmt.Printf(url1, branch, branch)
		fmt.Println("")
		fmt.Printf(url2, branch, branch)
		fmt.Println("")
		fmt.Printf(url3, branch, s, branch, s)
		fmt.Println("")
	}
	io.Copy(os.Stdout, file)
}
