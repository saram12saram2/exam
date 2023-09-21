package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 2 {
		n := 0
		for _, w := range os.Args[1] {
			if w < '0' || w > '9' {
				fmt.Printf("ERROR: cannot convert to roman digit\n")
				return
			}
			n = n*10 + int(w-'0')
		}
		if n == 0 || n >= 4000 {
			fmt.Printf("ERROR: cannot convert to roman digit\n")
			return
		}
		num := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
		romnum := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
		mathnum := []string{"M", "(M-C)", "D", "(D-C)", "C", "(C-X)", "L", "(L-X)", "X", "(X-I)", "V", "(V-I)", "I"}
		var str, sum string // str := "" sum := ""
		for i := range num {
			temp := n / num[i]
			for temp > 0 {
				n -= num[i]
				str += romnum[i]
				sum += mathnum[i] + "+"
				temp--
			}
		}
		fmt.Printf("%s\n%s\n", sum[:len(sum)-1], str)
	}
}

// Second Version
/*func main() {
	if len(os.Args) == 2 {
		n := Atoi(os.Args[1])
		if n == 123123 || n >= 4000 {
			fmt.Printf("ERROR: cannot convert to roman digit\n")
		} else {
			rn(n)
		}
	}
}
func Atoi(str string) int {
	neg := 1 // var neg bool // for false
	if str[0] == '-' {
		neg = -1 // neg = true // for true
		str = str[1:]
	}
	n := 0 // var n int
	for _, w := range str {
		if w < '0' || w > '9' { // if w < 48 || w > 58
			return 123123
		}
		n = n*10 + int(w-'0') // n*10 + int(w) - 48
	}
	return neg * n
}

func rn(n int) {
	num := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	romnum := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	mathnum := []string{"M", "(M-C)", "D", "(D-C)", "C", "(C-X)", "L", "(L-X)", "X", "(X-I)", "V", "(V-I)", "I"}

	var str, sum string // str := "" sum := ""
	for i := range num {
		temp := n / num[i]
		for temp > 0 {
			n -= num[i]
			str += romnum[i]
			sum += mathnum[i] + "+"
			temp--
		}
	}
	fmt.Printf("%s\n%s\n", sum[:len(sum)-1], str)
}*/

/*Third Variant
func main() {
	if len(os.Args) == 2 {
		a := Atoi(os.Args[1])
		if a == 123123 || a >= 4000 {
			fmt.Printf("ERROR: cannot convert to roman digit\n")
		} else {
			Roman(a)
		}
	}
}
func Atoi(s string) int {
	n := 0
	t := 1
	if s[0] == '-' {
		t = -1
		s = s[1:]
	}
	for _, w := range s {
		if w >= '0' && w <= '9' {
			n = n*10 + int(w-48)
		} else {
			return 123123
		}
	}
	return t * n
}
func Roman(n int) {
	num := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	rom_n := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	res_n := []string{"M", "(M-C)", "D", "(D-C)", "C", "(C-X)", "L", "(L-X)", "X", "(X-I)", "V", "(V-I)", "I"}
	i := 0
	var res, rom string
	for n > 0 {
		k := n / num[i]
		for j := 0; j < k; j++ {
			res += res_n[i] + "+"
			rom += rom_n[i]
			n -= num[i]
		}
		i++
	}
	res = res[:len(res)-1]
	fmt.Printf(res)
	fmt.Printf("\n")
	fmt.Printf(rom)
	fmt.Printf("\n")
}*/
