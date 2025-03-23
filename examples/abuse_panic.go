package main

import "fmt"

type Mountain struct{}

func (m *Mountain) Climb() {
	fmt.Println("climbing")
}

func ClimbAllPanicRecover(mountains []Mountain) {
	defer func() {
		fmt.Println("ops, went too high...")
		recover()
	}()
	for i := 0; ; i++ {
		mountains[i].Climb() // panics when i == len(mountains)
	}
	fmt.Println("unreachable, even after recover")
}

func main() {
	mountains := make([]Mountain, 5)
	ClimbAllPanicRecover(mountains)
	fmt.Println("at the peak!")
}
