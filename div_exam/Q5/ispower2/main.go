// package main

// import (
// 	"os"
// 	"strconv"

// 	"github.com/01-edu/z01"
// )

// func main() {
// 	if len(os.Args) == 2 {
// 		// This line uses the strconv.Atoi function to convert the command-line
// 		// argument from a string to an integer. It also ignores the second return value
// 		// from the function, which is an error message. If the conversion fails, the
// 		// program will panic.
// 		num, _ := strconv.Atoi(os.Args[1])

// 				// This condition checks whether the input number is a power of two.
// 		// If it is, the program will output "true". Otherwise, it will output "false".
// 		if Ispower2(num) {
// 			for _, w := range "true" {
// 				z01.PrintRune(w)
// 			}
// 		} else {
// 			for _, w := range "false" {
// 				z01.PrintRune(w)
// 			}
// 		}
// 		z01.PrintRune('\n')
// 	}
// }

// func Ispower2(n int) bool {
// 	return n != 0 && n&(n-1) == 0
// }

package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) == 2 {

		num, _ := strconv.Atoi(os.Args[1])

		if ispower(num) {
			for _, val := range "true" {
				z01.PrintRune(val)
			}
		} else {
			for _, val := range "false" {
				z01.PrintRune(val)
			}
		}
		z01.PrintRune('\n')
	}
}

func ispower(n int) bool {
	return n != 0 && n&(n-1) == 0
}

/* Second Variant
func main() {
	args := os.Args[1:]
	lenArgs := 0
	for range args {
		lenArgs++
	}
	if lenArgs != 1 {
		return
	}
	num, err := strconv.Atoi(args[0])
	if err != nil {
		return
	}
	status := "false"
	for i := 2; i <= num; i = i * 2 {
		if i == num {
			status = "true"
		}
	}
	for _, letter := range status {
		z01.PrintRune(letter)
	}
	z01.PrintRune('\n')
}*/
