package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) == 2 {
		for _, i := range os.Args[1] {
			if i >= 'A' && i <= 'Z' {
				i = ('Z' - (i - 'A'))
			} else if i >= 'a' && i <= 'z' {
				i = ('z' - (i - 'a'))
			}
			z01.PrintRune(i)
		}
		z01.PrintRune('\n')
	}
}
