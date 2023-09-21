package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) == 2 {
		Brainfuck(os.Args[1])
	} else {
		return
	}
}

func Brainfuck(s string) {
	b := [2048]byte{}
	x := 0
	p := &b[x]
	depth := 0
	for i := 0; i < len(s); i++ {
		c := s[i]
		switch c {
		case '>':
			x++
		case '<':
			x--
		case '-':
			*p--
		case '+':
			*p++
		case '.':
			Printer(string(*p))
		case '[':
			depth++
			if *p == 0 {
				oldDepth := depth - 1
				for oldDepth != depth {
					i++
					if s[i] == '[' {
						depth++
					} else if s[i] == ']' {
						depth--
					}
				}
			}
		case ']':
			depth--
			if *p != 0 {
				oldDepth := depth + 1
				for oldDepth != depth {
					i--
					if s[i] == '[' {
						depth++
					} else if s[i] == ']' {
						depth--
					}
				}
			}
		}
		p = &b[x]
	}
}

func Printer(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}
