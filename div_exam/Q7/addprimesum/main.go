package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) == 2 {
		n := 0 // var n int
		for i := 1; i <= Atoi(os.Args[1]); i++ {
			if Isprime(i) {
				n += i
			}
		}
		res := Itoa(n)
		for i := 0; i < len(res); i++ {
			z01.PrintRune(rune(res[i]))
		}
		// for _, i := range res {
		//	z01.PrintRune(i)
		//}
	} else {
		z01.PrintRune('0')
	}
	z01.PrintRune('\n')
}

func Isprime(nb int) bool {
	if nb <= 1 {
		return false
	}
	if nb > 1 {
		for i := 2; i <= nb/2; i++ {
			if nb%i == 0 {
				return false
			}
		}
	}
	return true
}

// func Isprime(nb int) bool {
// 	if nb > 1 {
// 		for i := 2; i <= nb/2; i++ {
// 			if nb%i == 0 {
// 				return false
// 			}
// 		}
// 	} else {
// 		return false
// 	}
// 	return true
// }

func Atoi(str string) int {
	neg := 1 // var neg bool // for false
	if str[0] == '-' {
		neg = -1 // neg = true // for true
		str = str[1:]
	}
	n := 0 // var n int
	for _, w := range str {
		if w < '0' || w > '9' { // if w < 48 || w > 58
			os.Exit(0)
		}
		n = n*10 + int(w-'0') // n*10 + int(w) - 48
	}
	return neg * n
}

func Itoa(n int) string {
	neg := ""
	if n < 0 {
		neg = "-"
		n = -n
	}
	if n == 0 {
		z01.PrintRune('0')
	}
	str := ""
	for n > 0 {
		str = string(rune(n%10+48)) + str
		n = n / 10
	}
	return neg + str
}
