
package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) == 3 {
		check := make(map[string]bool)
		for _, a := range os.Args[1] {
			for _, b := range os.Args[2] {
				if a == b {
					if _, value := check[string(a)]; !value {
						check[string(a)] = true
						z01.PrintRune(a)
					}
				}
			}
		}
	}
	z01.PrintRune('\n')
}

// package main

// import (
// 	"fmt"
// 	"os"
// 	"github.com/01-edu/z01"
// )

// func main() {
// 	if len(os.Args) != 3 {
// 		return
// 	}

// 	str1, str2 := os.Args[1], os.Args[2]
// 	var res string

// 	for _, r1 := range str1 {
// 		if !containsRune(str2, r1) && !containsRune(res, r1) {
// 			res = res + string(r1)
// 		}
// 	}

// 	fmt.Println(res)
// }

// func containsRune(s string, r rune) bool {
// 	for _, v := range s {
// 		if v == r {
// 			return true
// 		}
// 	}
// 	return false
// }


/*Second Variant
func check(s string, a rune) bool {
	for i := 0; i < len(s); i++ {
		if rune(s[i]) == a {
			return false
		}
	}
	return true
}

func main() {
	a := os.Args[1]
	b := os.Args[2]
	var str string
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(b); j++ {
			if a[i] == b[j] {
				if check(str, rune(a[i])) {
					str += string(a[i])
				}
			}
		}
	}
	// fmt.Println(string(s))
	for _, k := range str {
		z01.PrintRune(k)
	}
	z01.PrintRune('\n')
}*/
