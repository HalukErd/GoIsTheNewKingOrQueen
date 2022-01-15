package main

import (
	"fmt"
)

type person struct {
	fName string
	lName string
}

type secretAgent struct {
	person
	licenseToKill bool
}

type human interface {
	speak()
}

func main() {
	me := person{
		fName: "Haluk",
		lName: "Erd",
	}
	p007 := secretAgent{
		person{
			fName: "James",
			lName: "Bond",
		},
		true,
	}

	//me.speak()
	//p007.speak()
	saySth(me)
	saySth(p007.person)
	saySth(p007)
}

func (p person) speak() {
	fmt.Println(p.fName, p.lName, `says, "Good morning"`)
}
func (sa secretAgent) speak() {
	fmt.Printf(`%v %v says, "%v, %v %v'"`, sa.fName, sa.lName, sa.lName, sa.fName, sa.lName)
}

func saySth(h human) {
	h.speak()
}
