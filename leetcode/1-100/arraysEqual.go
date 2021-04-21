package main

import "fmt"

func main() {
	a := [2]int{5, 6}
	b := [2]int{5, 6}
	if a == b {
		fmt.Println("equal")
	} else {
		fmt.Println("not equal")
	}
	c := a[0:1]
	d := b[0:1]
	if c == d {
		fmt.Println("equal")
	} else {
		fmt.Println("not equal")
	}

}
