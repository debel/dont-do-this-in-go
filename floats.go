package main

import "fmt"

func smartSumEqual(u, v, w float64) bool {
	return (u+v == w) || (u == w-v) || (v == w-u)
}

func actualSolution(u, v, w float64) bool {
	return u+v-w < 1e-8
}

func main() {
	x, y, z := 0.1, 0.2, 0.3

	fmt.Println("0.1 + 0.2 == 0.3 ?", x+y == z)
	fmt.Println("'smart' sum equals ?", smartSumEqual(x, y, z))
	fmt.Println("actual sum equals ?", actualSolution(x, y, z))
}
