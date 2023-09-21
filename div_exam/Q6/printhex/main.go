package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) == 2 {
		num, err := strconv.Atoi(os.Args[1])
		if err != nil {
			Print("ERROR")
			return
		}
		hex := strconv.FormatInt(int64(num), 16)
		Print(hex)
	}
}

func Print(s string) {
	for _, str := range s {
		z01.PrintRune(str)
	}
	z01.PrintRune('\n')
}
