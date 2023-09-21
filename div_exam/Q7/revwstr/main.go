package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	arg := os.Args[1]
	count := 0
	arr := []string{}
	word := ""

	for i := 0; i < len(arg); i++ {
		if arg[i] == ' ' && count == 1 {
			arr = append(arr, word)
			count = 0
			word = ""
		}
		if arg[i] != ' ' {
			word += string(arg[i])
			count = 1
		}
	}

	if count != 0 {
		arr = append(arr, word)
	}

	for i := len(arr) - 1; i >= 0; i-- {
		for _, v := range arr[i] {
			z01.PrintRune(rune(v))
		}
		if i != 0 {
			z01.PrintRune(rune(32))
		}
	}
	z01.PrintRune('\n')
}

//Incorrect Variant
/*func main() {
	if os.Args[1] == "" {
		z01.PrintRune('\n')
		return
	}
	if len(os.Args) == 2 {
		new := ""
		old := os.Args[1] + " $"
		idx := 0
		for i := range old {
			if old[i] == ' ' && old[i+1] != ' ' {
				new = old[idx:i] + " " + new
				idx = i + 1
			}
		}
		for _, v := range new {
			z01.PrintRune(v)
		}
	}
	z01.PrintRune('\n')
}*/
