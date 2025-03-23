package main

import "fmt"

type aStruct struct {
	anInt   int
	aString string
}

func main() {
	var s aStruct

	fmt.Println(s)
}
