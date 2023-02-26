package main

import "fmt"

func main() {

	a := []int{1, 2, 3, 4, 5, 6}
	b := []int{0, 0, 0}
	fmt.Println(a)
	fmt.Println(b)

	copy(a[1:], b)
	fmt.Println(a)
	fmt.Println(b)

}
