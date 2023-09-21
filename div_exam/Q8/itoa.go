package piscine

func Itoa(n int) string {
	neg := ""
	if n < 0 {
		neg = "-"
		n = -n
	}
	if n == 0 {
		return "0"
	}
	str := ""
	for n > 0 {
		str = string(rune(n%10+48)) + str
		n = n / 10
	}
	return neg + str
}

// TEST the function
/*func main() {
	x := -12345
	y := Itoa(x)
	fmt.Println(y)
}*/

//Second Variant
/*func Itoa(n int) string {
	neg := ""
	if n < 0 {
		neg = "-"
		// neg = "" // if positive needed
		n = -n
	}
	if n == 0 {
		return string(rune('0')) // return string(rune(48))
	}
	str := ""
	for n > 0 {
		str = string(rune(n%10+48)) + str
		n = n / 10
	}
	return neg + str
}
*/
//Third Variant
/*func Itoa(n int) string {
	var a, b []byte
	var c, d int
	for {
		if n < 0 {
			n = -1 * n
			// b = append(b, '-') // UNCOMMENT if negative NEEDED
		}
		c = n / 10
		d = n - c*10
		n = c
		a = append(a, byte('0'+d))
		if n == 0 {
			break
		}
	}
	for i := 0; i < len(a); i++ {
		b = append(b, a[len(a)-i-1])
	}
	return string(b)
}*/

//Fouth Variant but it doesn't work on the exams
/*func Itoa(n int) string {
	neg := ""
	if n < 0 {
		neg = "" // neg = "-"    -> if negative NEEDED
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
}*/
