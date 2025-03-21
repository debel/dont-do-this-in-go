package main

import "fmt"

type point struct{ x, y float64 }

type union interface {
	string | int | bool | point
}

func use[T union](p T) {
	fmt.Println(p)
}

func main() {
	use("what?")
	use(5) // change to 5.1
	use(true)
	use(point{x: 3.14, y: 2.71})
}
