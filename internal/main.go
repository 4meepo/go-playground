package main

import "fmt"

func main() {
	a := make([]int, 0, 2)
	b := append(a, 1)
	c := append(a, 2)


	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	d := f(a)
	fmt.Println(a)
	fmt.Println(d)
	// print a b c
	// a: len 0 []
	// b: len 1 [2]
	// c: len 1 [2]
	// a []
	// d [3]

}
func f(a []int) []int{
	a = append(a, 3)
	return a
}
