package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 3 {
		return
	}
	defer z01.PrintRune('\n')
	count := 0
	for _, v := range os.Args[2] {
		if len(os.Args[1]) == count {
			z01.PrintRune('1')
			return
		}
		if v == rune(os.Args[1][count]) {
			count++
		}
	}
	z01.PrintRune('0')
}
