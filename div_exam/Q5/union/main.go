package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) == 3 {
		check := make(map[string]bool)
		for _, a := range os.Args[1] {
			if _, value := check[string(a)]; !value {
				check[string(a)] = true
				z01.PrintRune(a)
			}
		}
		for _, b := range os.Args[2] {
			if _, value := check[string(b)]; !value {
				check[string(b)] = true
				z01.PrintRune(b)
			}
		}
	}
	z01.PrintRune('\n')
}

//first var
// package main

// import (
// 	"fmt"
// 	"os"
// )

// func main() { 
// 	// Check that there are exactly 3 command line arguments (including program name)
// 	if len(os.Args) != 3 {
// 		fmt.Println()
// 		return
// 	}
// 	var res string

// 	// Loop over all the characters in the concatenation of the two command line arguments
// 	for _, val := range os.Args[1] + os.Args[2] {
// 		if !double(res, val) {
// 			res += string(val)
// 		}
// 	}
// 	fmt.Println(res)
// }
// // Loop over all the characters in the string s
// func double(s string, r rune) bool {
// 	for _, val := range s {
// 		if val == r {
// 			return true
// 		}
// 	}
// 	return false
// }

/*Second Variant
func main() {
	arg := os.Args[1:]
	if len(arg) != 2 {
		z01.PrintRune('\n')
	} else {
		sum := arg[0] + arg[1]
		result := ""
		for _, w := range sum {
			if !contain(result, w) {
				result += string(w)
			}
		}
		for _, w := range result {
			z01.PrintRune(w)
		}
		z01.PrintRune('\n')
	}
}

func contain(s string, r rune) bool {
	for _, w := range s {
		if w == r {
			return true
		}
	}
	return false
}*/
