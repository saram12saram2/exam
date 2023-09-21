package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	num, _ := strconv.Atoi(os.Args[1])
	var res string
	for i := 1; i < 10; i++ {
		res += strconv.Itoa(i) + " x " + strconv.Itoa(num) + " = " + strconv.Itoa(i*num) + "\n"
	}
	for _, r := range res {
		z01.PrintRune(r)
	}
}

// The Second variant - the shortest and the easiest one
/*func main() {
	if len(os.Args) == 2 {
		arg := os.Args[1:]
		nbr, _ := strconv.Atoi(arg[0])
		for i := 1; i <= 9; i++ {
			z01.PrintRune(rune(i) + 48)
			z01.PrintRune(' ')
			z01.PrintRune('x')
			z01.PrintRune(' ')
			for _, v := range arg[0] {
				z01.PrintRune(v)
			}
			z01.PrintRune(' ')
			z01.PrintRune('=')
			z01.PrintRune(' ')
			res := strconv.Itoa(i * nbr)
			for _, v := range res {
				z01.PrintRune(v)
			}
			z01.PrintRune('\n')
		}
	}
} */

//The Third variant
/*func main() {
	if len(os.Args) != 2 {
		z01.PrintRune('\n')
		return
	}
	arg, _ := strconv.Atoi(os.Args[1])
	for i := 1; i <= 9; i++ {
		Print(Itoa(i) + " " + "x" + " " + Itoa(arg) + " " + "=" + " " + Itoa(arg*i))
	}
}
func Print(str string) {
	for _, w := range str {
		z01.PrintRune(w)
	}
	z01.PrintRune('\n')
}
func Itoa(n int) string {
	var flag bool
	a := ""
	if n < 0 {
		n *= -1
		flag = true
	}
	for n != 0 {
		a = string(rune(n%10)+48) + a
		n = n / 10
	}
	if flag {
		a = "-" + a
	}
	return a
}*/
