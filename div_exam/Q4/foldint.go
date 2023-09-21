package main

import (
	"fmt"

	"github.com/01-edu/z01"
)

func main() {
	table := []int{1, 2, 3}
	ac := 93
	FoldInt(Add, table, ac)
	FoldInt(Mul, table, ac)
	FoldInt(Sub, table, ac)
	fmt.Println()

	table = []int{0}
	FoldInt(Add, table, ac)
	FoldInt(Mul, table, ac)
	FoldInt(Sub, table, ac)
}

func Add(acc, cur int) int {
	return acc + cur
}

func Mul(acc, cur int) int {
	return acc * cur
}

func Sub(acc, cur int) int {
	return acc - cur
}

// FoldInt is a function that applies an arithmetic function on each element of a slice of integers and an initial value n,
// then prints the result of the final calculation.

func FoldInt(f func(int, int) int, a []int, n int) {
	// Iterate over the slice and apply the arithmetic function to each element and the previous result, storing the new result in n.
	for i := 0; i < len(a); i++ {
		n = f(n, a[i])
	}
	// Convert the final result to a string and print it using the z01 package.
	for _, j := range Itoa(n) {
		z01.PrintRune(j)
	}
	z01.PrintRune('\n')
}

// Itoa is a function that converts an integer to a string.
func Itoa(n int) string {
	neg := ""  // initialize a string to store a negative sign if needed
	if n < 0 { // if the number is negative, set neg to "-" and make n positive
		neg = "-"
		n = -n
	}
	if n == 0 { // if the number is 0, print a 0 character using the z01 package
		z01.PrintRune('0')
	}
	str := ""   // initialize an empty string to store the digits of the number
	for n > 0 { // while n is greater than 0
		str = string(rune(n%10+48)) + str // get the rightmost digit of n, convert it to a character, and add it to the left of str
		n = n / 10                        // remove the rightmost digit from n
	}
	return neg + str // add the negative sign (if needed) to the left of str and return the resulting string
}


/*Second Variant of Itoa
func Itoa(n int) string {
	ch := ""
	if n < 0 {
		n = -n
		ch = "-"
	}
	num := "0123456789"
	if n < 10 {
		return ch + num[n:n+1]
	}
	return ch + Itoa(n/10) + num[n%10:n%10+1]
}*/
