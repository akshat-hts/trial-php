package main

import (
	"fmt"
)

type Dog struct {}	// created custom datatype
type Cat struct {}	

func (d Dog) Speak() string {	// attached methods to those structs	
	return "Woof"
}	

func (c Cat) Speak() string { 	// attached methods to data-type
	return "Meow"
}

type Speaker interface {		// interface simply defines the behavior
	Speak() string
}
func makeSound(s Speaker) {
	fmt.Println(s.Speak())
}

func main() {
	d := Dog{}
	c := Cat{}

	makeSound(d)
	makeSound(c)
}