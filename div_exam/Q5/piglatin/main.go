package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) == 2 {
		res := ""
		for i, w := range os.Args[1] {
			if w == 'a' || w == 'e' || w == 'i' || w == 'o' || w == 'u' || w == 'A' || w == 'E' || w == 'I' || w == 'O' || w == 'U' {
				res = os.Args[1][i:] + os.Args[1][:i] + "ay"
				break
			} else {
				res = "No vowels"
			}
		}
		if res == "" {
			return
		}
		for _, w := range res {
			z01.PrintRune(w)
		}
		z01.PrintRune(10)
	}
}

// func main() {
// 	var res string
// 	if len(os.Args) != 2 { // || os.Args[1] == "" - if needed in the exam
// 		return
// 	}
// 	arr := os.Args[1]
// 	for i, w := range arr {
// 		if w == 'a' || w == 'e' || w == 'i' || w == 'o' || w == 'u' || w == 'A' || w == 'E' || w == 'I' || w == 'O' || w == 'U' {
// 			res = arr[i:] + arr[:i] + "ay"
// 			break
// 		} else {
// 			res = "No vowels"
// 		}
// 	}
// 	if res == "" {
// 		return
// 	}
// 	for _, w := range res {
// 		z01.PrintRune(w)
// 	}
// 	z01.PrintRune(10)
// }
