package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args[1:]) > 9 {
		z01.PrintRune(rune((len(os.Args[1:]) / 10) + 48))
		z01.PrintRune(rune((len(os.Args[1:]) % 10) + 48))
	} else {
		z01.PrintRune(rune(len(os.Args[1:])) + 48)
	}
	z01.PrintRune(10)
}

//Second Variant
/*func main() {
    if len(os.Args) > 1 {
        n := len(os.Args[1:])
        for _, k := range strconv.Itoa(n) {
            z01.PrintRune(k)
        }
        z01.PrintRune('\n')
    } else {
        z01.PrintRune('0')
        z01.PrintRune('\n')
        return
    }
}/*

// Third Variant
/* func main() {
	if len(os.Args[1:]) > 9 {
		num1 := len(os.Args[1:]) / 10
		num2 := len(os.Args[1:]) % 10
		z01.PrintRune(rune(num1) + 48)
		z01.PrintRune(rune(num2) + 48)
	} else {
		z01.PrintRune(rune(len(os.Args[1:])) + 48)
	}
	z01.PrintRune(10)
} /*

// Fourth Variant
/*func main() {
	Printnbr(len(os.Args[1:]))
}

func Printnbr(n int) {
	//var flag bool
	//if n == 0 {
	//	z01.PrintRune('0')
	//}
	//if n < 0 {
	//	flag = true
	//	n = -n
	//}

	var temp []rune
	for n != 0 {
		temp = append(temp, rune(n%10+'0'))
		n /= 10
	}

	for i := len(temp) - 1; i >= 0; i-- {
		//if flag {
		//	z01.PrintRune('-')
		//	flag = false
		//}
		//}
		z01.PrintRune(temp[i])
	}
	z01.PrintRune('\n')
}*/

// Fifth Variant
/*func main() {
	count := 0
	for _, args := range os.Args[1:] {
		if args != "\n" {
			count++
		}
	}
	PrintCount(count)
}

func PrintCount(count int) {
	var arr string = ""
	for count > 0 {
		arr = string(count%10+48) + arr
		count /= 10
	}
	for _, elem := range arr {
		z01.PrintRune((elem))
	}
	z01.PrintRune('\n')
}*/

// Sixth Variant
/*func main() {
    if len(os.Args) == 1 {
        z01.PrintRune('0')
        z01.PrintRune('\n')
        return
    }
    var str string
    var array []string
    arg := os.Args[1:]
    for , i := range arg {
        str += string(i)
        array = append(array, str)
    }
    for , j := range Itoa(len(array)) {
        z01.PrintRune(j)
    }
    z01.PrintRune('\n')
}

func Itoa(n int) string {
    var neg string
    if n < 0 {
        neg = "-"
        n = -n
    }
    if n == 0 {
        z01.PrintRune('0')
    }
    var str string
    for n > 0 {
        // str = string(rune(n%10 + 48))
        // n /= 10
        str = string(rune(n%10)+48) + str
        n /= 10

    }
    return neg + str
}*/
