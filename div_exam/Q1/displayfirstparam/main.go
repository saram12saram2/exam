package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) > 1 {
		for _, i := range os.Args[1] {
			z01.PrintRune(i)
		}
		z01.PrintRune('\n')
	}
}

// Second Variant
/*func main() {
	if len(os.Args) == 1 {
		return
	}
	for _, i := range os.Args[1] {
		z01.PrintRune(i)
	}
	z01.PrintRune('\n')
}*/
