package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) == 2 {
		res := ""
		num, _ := strconv.Atoi(os.Args[1])
		if num == 0 {
			res = "00000000"
		}
		for len(res) != 8 {
			res = string(rune(num%2+48)) + res
			num /= 2
		}
		for _, j := range res {
			z01.PrintRune(j)
		}
		z01.PrintRune('\n')
	}
}

// Second Variant
/*func main() {
	if len(os.Args) != 2 {
		return
	}
	result, err := strconv.Atoi(os.Args[1])
	if err != nil {
		for _, w := range "00000000" {
			z01.PrintRune(w)
		}
		z01.PrintRune('\n')
		return
	}
	slice := make([]int, 8)
	i := len(slice) - 1
	for result != 0 {
		slice[i] = result % 2
		result /= 2
		i--
	}
	for _, w := range slice {
		z01.PrintRune(rune(w + 48))
	}
	z01.PrintRune('\n')
}/*

// Third Variant
/*func main() {
	if len(os.Args) != 2 {
		return
	}
	a, _ := strconv.Atoi(os.Args[1])
	var str string
	if a == 0 {
		str = "00000000"
	}
	for len(str) != 8 {
		str = string(rune(a%2+48)) + str
		a /= 2
	}
	str2 := str[4:8] + str[0:4]
	for _, i := range str {
		z01.PrintRune(i)
	}
	z01.PrintRune('\n')
	for _, i := range str2 {
		z01.PrintRune(i)
	}
	z01.PrintRune('\n')
}*/
