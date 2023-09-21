package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) == 2 {
		if i, err := strconv.Atoi(os.Args[1]); err == nil {
			Fprime(i)
		}
	}
}

func Fprime(nb int) {
	if nb <= 1 {
		return
	}
	i := 2
	for nb > 1 {
		if nb%i == 0 {
			fmt.Print(i)
			nb = nb / i
			if nb > 1 {
				fmt.Print("*")
			}
			i--
		}
		i++
	}
	fmt.Println()
}
