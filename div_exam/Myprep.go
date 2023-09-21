// displayfirstparam

func main() {
	if len(os.Args) > 1 {
		for _, i := range os.Args[1] {
			z01.PrintRune(i)
		}
		z01.PrintRune('\n')
	}
}

//displaylastparam

func main() {
	if len(os.Args) > 1 {
		for _, i := range os.Args[len(os.Args)-1] {
			z01.PrintRune(i)
		}
		z01.PrintRune('\n')
	}
}

// paramcount

func main() {
	if len(os.Args[1:]) > 9 {
		z01.PrintRune(rune((len(os.Args[1:]) / 10) + 48))
		z01.PrintRune(rune((len(os.Args[1:]) % 10) + 48))
	} else {
		z01.PrintRune(rune(len(os.Args[1:])) + 48)
	}
	z01.PrintRune(10)
}

//////////////////////////////////////////////////////////// QUEST2

// strlen

func StrLen(s string) int {
	return len([]rune(s))
}

// max

func Max(a []int) int {
	max := a[0]
	for _, val := range a {
		if val > max {
			max = val
		}
	}
	return max
}

// wdmatch

func main() {
	if len(os.Args) == 3 {
		arg := os.Args[1:]
		for i, j := 0, 0; i < len(arg[0]) && j < len(arg[1]); {
			if arg[0][i] == arg[1][j] {
				i++
				if i == len(arg[0]) {
					for _, w := range arg[0] {
						z01.PrintRune(w)
					}
					z01.PrintRune('\n')
				}
			}
			j++
			if j == len(arg[1]) {
				return
			}
		}
	}
}

//////////////////////////////////////////////////////////// QUEST3

func FirstRune(s string) rune {
	return []rune(s)[0]
}

func LastRune(s string) rune {
	return []rune(s)[len(s)-1]
}

// rot13

func main() {
	if len(os.Args) == 2 {
		for _, i := range os.Args[1] {
			if i >= 'A' && i <= 'M' || i >= 'a' && i <= 'm' {
				i += 13
			} else if i >= 'N' && i <= 'Z' || i >= 'n' && i <= 'z' {
				i -= 13
			}
			z01.PrintRune(i)
		}
		z01.PrintRune('\n')
	}
}

// lastword

func main() {
	if len(os.Args) == 2 {
		args := os.Args[1]
		res := "" // var res string
		for i := len(args) - 1; i >= 0; i-- {
			if args[i] != ' ' {
				res = string(args[i]) + res
			}
			if args[i] == ' ' && res != "" {
				break
			}
		}
		if res == "" {
			return
		}
		for _, v := range res {
			z01.PrintRune(v)
		}
		z01.PrintRune('\n')
	}
}

// reduceint

func ReduceInt(a []int, f func(int, int) int) {
	n := a[0]

	for i := 1; i < len(a); i++ {
		n = f(n, a[i])
	}

	for _, val := range strconv.Itoa(n) {
		z01.PrintRune(val)
	}
	z01.PrintRune('\n')
}

////////////////////////////////////////////////////////// QUEST 4

func Chunk(slice []int, size int) {
	bs := [][]int{}
	if len(slice) == 0 {
		fmt.Println() // fmt.Println(bs)
		return
	}
	if size <= 0 {
		fmt.Println()
		return
	}

	for len(slice) > size {
		bs = append(bs, slice[:size])
		slice = slice[size:]
	}

	bs = append(bs, slice)
	fmt.Println(bs)
}

// searchreplace

func main() {
	if len(os.Args) == 4 {
		for _, v := range os.Args[1] {
			if v == []rune(os.Args[2])[0] {
				v = []rune(os.Args[3])[0]
			}
			z01.PrintRune(v)
		}
		z01.PrintRune('\n')
	}
}

// tabmult

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

// alphamirror

func main() {
	if len(os.Args) == 2 {
		for _, i := range os.Args[1] {
			if i >= 'A' && i <= 'Z' {
				i = ('Z' - (i - 'A'))
			} else if i >= 'a' && i <= 'z' {
				i = ('z' - (i - 'a'))
			}
			z01.PrintRune(i)
		}
		z01.PrintRune('\n')
	}
}

// compare
func Compare(a, b string) int {
	if a == b {
		return 0
	}

	if a > b { // add the negative sign (if needed) to the left of str and return the resulting string
		return 1
	}

	if a < b {
		return -1
	}
	return 0
}

// doop

func main() {
	if len(os.Args) == 4 {

		args := os.Args[1:]
		a := Atoi(args[0])
		b := Atoi(args[2])

		if a > 0 && b > 0 && (a+b) < 0 || a < 0 && b < 0 && (a+b) > 0 {
			return
		}
		if len(args[0]) >= 19 {
			return
		}
		switch args[1] {
		case "+":
			fmt.Printf("%d\n", a+b)
		case "-":
			fmt.Printf("%d\n", a-b)
		case "/":
			if b == 0 {
				fmt.Printf("No division by 0\n")
				return
			}
			fmt.Printf("%d\n", a/b)
		case "*":
			fmt.Printf("%d\n", a*b)
		case "%":
			if b == 0 {
				fmt.Printf("No modulo by 0\n")
				return
			}
			fmt.Printf("%d\n", a%b)
		}
	}
}

func Atoi(str string) int {
	neg := 1
	if str[0] == '-' {
		neg = -1
		str = str[1:]
	}
	n := 0

	for _, val := range str {
		if val < '0' || val > '9' {
			os.Exit(0)
		}

		n = n*10 + int(val-'0')
	}
	return neg * n
}

// findprevprime

func FindPrevPrime(nb int) int {
	if nb <= 1 {
		return 0
	}
	for nb > 1 {
		if isprime(nb) {
			break
		}
		nb--
	}
	return nb
}

func isprime(nb int) bool {
	for i := 2; i <= nb/2; i++ {
		if nb%i == 0 {
			return false
		}
	}
	return true
}

// foldint

func FoldInt(f func(int, int) int, a []int, n int) {
	for i := 0; i < len(a); i++ {
		n = f(n, a[i])
	}
	for _, j := range Itoa(n) {
		z01.PrintRune(j)
	}
	z01.PrintRune('\n')
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

////////////////////////////////////////////////////////// QUEST 5

// romannumbers

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

// ispowerof2

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

// union

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

// inter

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

// printbits

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

// swapbits

func SwapBits(octet byte) byte {
	return octet<<4 + octet>>4
}

// reversebits
// Reversebits(00100110) -> 01100100

func ReverseBits(oct byte) byte {
	var res byte
	for i := 8; i > 0; i-- {
		res = (res << 1) | (oct & 1)
		oct >>= 1
	}
	return res
}

// piglatin

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
