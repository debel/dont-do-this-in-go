package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4}
	b := a[:2] /* [1, 2] */
	c := a[2:] /* [3, 4] */

	b = append(b, 5) // HL

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
