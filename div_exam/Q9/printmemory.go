package piscine

import (
	"unicode"

	"github.com/01-edu/z01"
)

func PrintMemory(arr [10]byte) {
	var temp [10]string
	for i := 0; i < 10; i++ {
		temp[i] = H(int(arr[i]))
	}
	for i, v := range temp {
		i++
		if ((i % 4) == 0) || i == len(temp) {
			Prln(v)
		} else {
			Pr(v + " ")
		}
	}
	for i, v := range arr {
		if unicode.IsGraphic(rune(v)) {
			Pr(string(v))
		} else {
			if i == len(arr)-1 {
				Prln(".")
				break
			}
			Pr(".")
		}
	}
}

func H(num int) (res string) {
	if num == 0 {
		return "00"
	}
	hex := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "a", "b", "c", "d", "e", "f"}
	for num > 0 {
		res = hex[num%16] + res
		num /= 16
	}
	return
}

func Prln(s string) {
	for _, v := range s {
		z01.PrintRune(v)
	}
	z01.PrintRune('\n')
}

func Pr(s string) {
	for _, v := range s {
		z01.PrintRune(v)
	}
}

//func main() {
//	PrintMemory([10]byte{'h', 'e', 'l', 'l', 'o', 16, 21, '*'})
//} - TEST FUNCTION
