package main

import (
	"fmt"
	"sync"
)

type SharedInt struct {
	mu  sync.Mutex
	val int
}

func (si *SharedInt) Val() int {
	si.mu.Lock()
	defer si.mu.Unlock() // 🔒
	return si.val        // 🔒
}

func (si *SharedInt) SetVal(val int) {
	si.mu.Lock()
	defer si.mu.Unlock() // 🔒
	si.val = val         // 🔒
}

func main() {
	myInt := &SharedInt{}

	wg := sync.WaitGroup{}
	wg.Add(200)

	for i := 0; i < 200; i += 1 {
		go func() { myInt.SetVal(myInt.Val() + 1); wg.Done() }()
	}

	wg.Wait()
	fmt.Println(myInt.val)
}
