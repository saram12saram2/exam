package main

import (
	"github.com/01-edu/z01"
)

func main() {
	a := "Hello World!"
	for _, hello := range a {
		z01.PrintRune(hello)
	}
	z01.PrintRune('\n')
}
