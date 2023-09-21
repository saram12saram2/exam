package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) == 2 {
		for _, i := range os.Args[1] {
			if i >= 'A' && i <= 'M' || i >= 'a' && i <= 'm' {
				i += 13
			} else if i >= 'N' && i <= 'Z' || i >= 'n' && i <= 'z' {
				i -= 13
			}
			z01.PrintRune(i)
		}
		z01.PrintRune('\n')
	}
}
