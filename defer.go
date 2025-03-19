package main

import "fmt"

func a(from string) {
	fmt.Println(from, ": a")
}

func b(from string) {
	fmt.Println(from, ": b")
}

func c(from string) {
	fmt.Println(from, ": c")
}

func x1() {
	defer a("x1")
	defer b("x1")
	c("x1")
}

func x2() {
	defer a("x2"); b("x2") // HLxxx
	c("x2")
}

func main() {
	x1()
	x2()
}
