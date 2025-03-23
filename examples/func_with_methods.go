package main

import "fmt"

type ImAStructTrustMe func()

func (ImAStructTrustMe) DoStuff() {
	fmt.Println("doing")
}

func (ImAStructTrustMe) AnotherMethod() {
	fmt.Println("much lol")
}

var myStruct ImAStructTrustMe = func() {
	fmt.Println("WTF?!")
}

func main() {
	myStruct.DoStuff()
	myStruct.AnotherMethod()

	// WTF!?
	myStruct() // HLxxx
}
