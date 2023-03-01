package main

import "fmt"

func main() {
	var a uint8 = 255
	var b uint8 = 1
	fmt.Printf("%d的二进制: %08b\n", a, a)
	fmt.Printf("  %d的二进制: %08b\n", b, b)
	fmt.Printf("a+b的二进制: %08b\n", a+b)
}
