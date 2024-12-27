package main

import (
	"fmt"
)

var act = true
var total int

func do() {
	act = true
}

func dont() {
	act = false
}

func mul(x, y int) {
	if act {
		total += x*y
	} 
}

func main() {
	total = 0
	first()
	fmt.Printf("first: %d\n", total)
	total = 0
	second()
	fmt.Printf("second: %d\n", total)
}