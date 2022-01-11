package main

import (
	"fmt"
	"log"
)

type contactInfo struct {
	email       string
	phoneNumber int
}
type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {

	me := person{
		firstName: "haluk",
		lastName:  "erd",
		contactInfo: contactInfo{
			email:       "erd.haluk@gmail.com",
			phoneNumber: 5454555454,
		},
	}
	//mePointer := &me
	//mePointer.updateName("Luciano")
	me.updateName("Luciano")
	me.print()
	fmt.Printf("%+v", *&me)
}

func (pointerToPerson *person) updateName(newName string) {
	(*pointerToPerson).firstName = newName
}

func (pointerToPerson person) print() {
	log.Printf("%+v", pointerToPerson)
}
