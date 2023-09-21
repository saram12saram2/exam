package main

import (
	"os"

	"github.com/01-edu/z01"
)

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

//Second Variant
/*func main() {
    if len(os.Args) != 3 {
        return
    }
    arg1 := os.Args[1]
    arg2 := os.Args[2]

    for i, j := 0, 0; i < len(arg1) && j < len(arg2); j++ {
        if j == len(arg2) {
            return
        }
        if arg1[i] == arg2[j] {
            i++
            if i == len(arg1) {
                for _, v := range arg1 {
                    z01.PrintRune(v)
                }
                z01.PrintRune(10)
            }
        }
    }
}/*

//Third Variant
/*func main() {
	arg := os.Args[1:]
	if len(arg) != 2 {
		return
	}
	runes1 := []rune(arg[0])
	runes2 := []rune(arg[1])
	counter := 0
	k := 0
	for i := 0; i < len(runes1); i++ {
		for j := k; j < len(runes2); j++ {
			if runes1[i] == runes2[j] {
				k = j + 1
				counter++
				break
			}
		}
	}
	if counter != len(runes1) {
		return
	} else {
		for _, s := range arg[0] {
			z01.PrintRune(s)
		}
	}
	z01.PrintRune('\n')
}*/

//Fouth Variant
/*func main() {
	args := os.Args
	if len(args) == 3 {
		word1 := args[1]
		word2 := args[2]
		i := 0
		for _, val := range word2 {
			if i == len(word1) {
				break
			}
			if rune(word1[i]) == val {
				i++
			}
		}
		if i == len(word1) {
			for _, val := range word1 {
				z01.PrintRune(val)
			}
			z01.PrintRune('\n')
		}
	}
}*/
