package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) == 2 {
		var a, b, c string
		os.Args[1] += " "
		for _, w := range os.Args[1] {
			if w == ' ' && b != "" {
				if a == "" {
					a = b
					b = ""
					continue
				}
				if c == "" {
					c = b
				} else {
					c = c + string(" ") + b
				}
				b = ""
			} else if w != ' ' {
				b = b + string(w)
			}
		}
		if c == "" {
			c = a
		} else {
			c = c + string(" ") + a
		}
		for _, k := range c {
			z01.PrintRune(k)
		}
	}
	z01.PrintRune('\n')
}
