package main

import (
	"strconv"

	"github.com/01-edu/z01"
)

func main() {
	mul := func(acc int, cur int) int {
		return acc * cur
	}
	sum := func(acc int, cur int) int {
		return acc + cur
	}
	div := func(acc int, cur int) int {
		return acc / cur
	}
	as := []int{500, 2}
	ReduceInt(as, mul)
	ReduceInt(as, sum)
	ReduceInt(as, div)
}

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
