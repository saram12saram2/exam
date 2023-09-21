package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		for _, w := range os.Args[1:] {
			fmt.Println(Brackets(w))
		}
	} else {
		return // fmt.Println() - for empty line
	}
}

func Brackets(s string) string {
	br := ""
	for _, r := range s {
		if (r == ')' || r == '}' || r == ']') && len(br) == 0 {
			return "Error"
		}
		if r == '(' || r == '{' || r == '[' {
			br += string(r)
		}
		if r == ')' && br[len(br)-1] == '(' {
			br = br[:len(br)-1]
		}
		if r == '}' && br[len(br)-1] == '{' {
			br = br[:len(br)-1]
		}
		if r == ']' && br[len(br)-1] == '[' {
			br = br[:len(br)-1]
		}
	}
	if len(br) == 0 {
		return "OK"
	}
	return "Error"
}
