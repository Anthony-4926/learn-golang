package main

import (
	"fmt"
	"strconv"
)

func main() {

	a, _ := strconv.ParseInt("01101100100000", 2, 64)
	b, _ := strconv.ParseInt("00000000100000", 2, 64)

	c := a & (-a)
	fmt.Println(a)
	fmt.Printf("b=%d, c=%d", b, c)

}
