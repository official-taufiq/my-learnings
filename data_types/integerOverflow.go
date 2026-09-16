package main

import (
	"fmt"
	"math"
)

func main() {
	inc32(math.MaxInt32)
	incUint(math.MaxUint)
	addInt(1, math.MaxInt)
}

// Detecting integer overflow during incrementing
func inc32(counter int32) {
	if counter == math.MaxInt32 { // compares with math.MaxInt32
		panic("int32 overflow")
	}
	fmt.Println(counter + 1)
}

func incUint(counter uint) {
	if counter == math.MaxUint { //compares with math.MaxUint
		panic("uint overflow")
	}
	fmt.Println(counter + 1)
}

// Detecting integer overflows during addition
func addInt(a, b int) {
	if a > math.MaxInt-b { // checks for possible overflow
		panic("int overflow")
	}
	fmt.Println(a + b)
}

// Detecting integer overflow during multiplication

func MultiplyInt(a, b int) int {
	if a == 0 || b == 0 { //If one of the operands is equal to 0, it directly returns 0.
		return 0
	}
	result := a * b
	if a == 1 || b == 1 { // Checks if one of the operands is equal to 1
		return result
	}
	if a == math.MinInt || b == math.MinInt { //Checks if one of the operands is equal to math.MinInt
		panic("integer overflow")
	}
	if result/b != a { // Checks if the multiplication leads to an integer overflow
		panic("integer overflow")
	}
	return result
}
