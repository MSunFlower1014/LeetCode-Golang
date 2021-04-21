package main

import "fmt"

func main() {
	var a, b int
	b = incr(a)
	fmt.Print(b)

	var c, d int
	d = incr2(c)
	fmt.Print(d)
}

func incr(a int) int {
	var b int
	defer func() {
		a++
		b++
	}()
	a++
	b = a
	return b
}

func incr2(a int) (b int) {
	defer func() {
		a++
		b++
	}()
	a++
	b = a
	return b
}
