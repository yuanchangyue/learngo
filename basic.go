package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

const FILENAME = "123.text"

func eular() {
	a := 3 + 4i
	fmt.Println(cmplx.Abs(a))
}

func enums() {
	const (
		java = iota
		_
		python
		golang
		js
	)
	// << 乘2的多少次方
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(java, python, golang)
	fmt.Println(b, kb, mb, gb, tb, pb)
}

func triangle() {
	a, b := 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(FILENAME, c)
}

func main() {
	eular()
	triangle()
	enums()
}
