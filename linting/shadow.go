package main

import "fmt"

func main() {
	i := 0

	if true {
		i := 1
		fmt.Println(i) //prints 1
	}
	fmt.Println(i) // but this prints 0
}
