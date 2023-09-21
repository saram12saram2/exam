package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) == 2 {
		args := os.Args[1]
		res := "" // var res string
		for i := len(args) - 1; i >= 0; i-- {
			if args[i] != ' ' {
				res = string(args[i]) + res
			}
			if args[i] == ' ' && res != "" {
				break
			}
		}
		if res == "" {
			return
		}
		for _, v := range res {
			z01.PrintRune(v)
		}
		z01.PrintRune('\n')
	}
}

// my code ? strings? func main() {
// 	if len(os.Args) > 1 {

// 		args1 := os.Args[1]

// 		words := strings.Fields(args1)
// 		lastw := words[len(words)-1]

// 		for _, val := range lastw {
// 			z01.PrintRune(val)
// 		}
// 		z01.PrintRune('\n')
// 	}
// }


/* Second variant
func main() {
	if len(os.Args) == 2 {
		sentence := os.Args[1]
		word := ""
		var result []string
		for i, letter := range sentence {
			if letter != ' ' {
				word += string(letter)
			} else if letter == ' ' && len(word) != 0 {
				result = append(result, word)
				word = ""
			}
			if i == len(sentence)-1 && len(word) != 0 {
				result = append(result, word)
			}
		}
		if len(result) > 0 {
			for _, letter := range result[len(result)-1] {
				z01.PrintRune(letter)
			}
			z01.PrintRune('\n')
		}
	}
} */
