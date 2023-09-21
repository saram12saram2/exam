// package main

// import (
// 	"fmt"
// 	"os"
// )

// func main() {
// 	if len(os.Args) == 4 { // Check if there are exactly 4 arguments
// 		arg := os.Args[1:] // Get the arguments
// 		a := Atoi(arg[0])  // Convert the first argument to an integer
// 		b := Atoi(arg[2])  // Convert the third argument to an integer
// 		// Check for overflow conditions
// 		if a > 0 && b > 0 && (a+b) < 0 || a < 0 && b < 0 && (a+b) > 0 {
// 			return // Exit the program if there's an overflow
// 		}
// 		if len(arg[0]) >= 19 { // Check if the length of the first argument is too big
// 			return // Exit the program if the length of the first argument is too big
// 		}
// 		switch arg[1] { // Check the operator
// 		case "+":
// 			fmt.Printf("%d\n", a+b) // Print the sum
// 		case "-":
// 			fmt.Printf("%d\n", a-b) // Print the difference
// 		case "*":
// 			fmt.Printf("%d\n", a*b) // Print the product
// 		case "/":
// 			if b == 0 {
// 				fmt.Printf("No division by 0\n") // Print an error message if the divisor is 0
// 				return // Exit the program
// 			}
// 			fmt.Printf("%d\n", a/b) // Print the quotient
// 		case "%":
// 			if b == 0 {
// 				fmt.Printf("No modulo by 0\n") // Print an error message if the divisor is 0
// 				return // Exit the program
// 			}
// 			fmt.Printf("%d\n", a%b) // Print the remainder
// 		}
// 	}
// }

// func Atoi(str string) int {
// 	neg := 1 // Initialize the sign to positive
// 	if str[0] == '-' {
// 		neg = -1 // Change the sign to negative if the string starts with a '-'
// 		str = str[1:] // Remove the '-' sign
// 	}
// 	n := 0 // Initialize the integer to 0
// 	for _, w := range str {
// 		if w < '0' || w > '9' { // Check if the character is not a digit
// 			os.Exit(0) // Exit the program if the character is not a digit
// 		}
// 		n = n*10 + int(w-'0') // Convert the character to an integer and add it to the total
// 	}
// 	return neg * n // Return the integer with the correct sign
// }

package main

import (
	"fmt"
	"os"
)

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

/* for example: 123 ->
n = 0 -> n = 0*10 + 1 = 1
n = 1 -> n = 1*10 + 2 = 12
n = 12 -> n = 12*10 + 3 = 123 */

//Second Variant
/*func main() {
    if len(os.Args) != 4 {
        return
    }
    v1 := Atoi(os.Args[1])
    op := os.Args[2]
    v2 := Atoi(os.Args[3])
    if v1 > 21000000000  v2 > 21000000000  v1 < -21000000000  v2 < -21000000000 {
        return
    }
    switch op {
    case "+":
        fmt.Printf("%v\n", v1+v2)
    case "-":
        fmt.Printf("%v\n", v1-v2)
    case "":
        fmt.Printf("%v\n", v1v2)
    case "/":
        if v2 == 0 {
            fmt.Printf("%v\n", "No division by 0")
        } else {
            fmt.Printf("%v\n", v1/v2)
        }
    case "%":
        if v2 == 0 {
            fmt.Printf("%v\n", "No modulo by 0")
        } else {
            fmt.Printf("%v\n", v1%v2)
        }
    }
}
func Atoi(s string) int {
    n := 0
    sign := 1
    if s[0] == '-' {
        sign = -1
        s = s[1:]
    }
    for _, v := range s {
        if v < '0'  v > '9' {
            os.Exit(0)
        }
        n = n10 + int(v-48)
    }
    return sign n
}*/

//Third Variat
/*func main() {
	if len(os.Args) != 4 {
		return
	}

	for _, s := range os.Args[1] {
		if s < '0' || s > '9' {
			return
		}
	}

	arg1 := Atoi(os.Args[1])
	arg2 := Atoi(os.Args[3])

	a := []rune(os.Args[1])
	b := []rune(os.Args[3])

	sign := os.Args[2]

	if arg1 < 0 {
		return
	} else if len(a) > 18 || len(b) > 18 {
		return
	}

	switch sign {
	case "+":
		res := arg1 + arg2
		fmt.Printf("%d\n", res)
	case "-":
		res := arg1 - arg2
		fmt.Printf("%d\n", res)
	case "*":
		res := arg1 * arg2
		fmt.Printf("%d\n", res)
	case "/":
		if arg2 == 0 {
			fmt.Printf("No division by 0\n")
			return
		}
		res := arg1 / arg2
		fmt.Printf("%d\n", res)
	case "%":
		if arg2 == 0 {
			fmt.Printf("No modulo by 0\n")
			return
		}
		res := arg1 % arg2
		fmt.Printf("%d\n", res)
	}
}

func Atoi(s string) int {
	var nbr int
	var negative bool

	if s[0] == '-' {
		negative = true
		s = s[1:]
	}

	for _, v := range s {
		if v < 48 && v > 58 {
			os.Exit(0)
		}
		nbr = nbr*10 + int(v) - 48
	}
	if negative {
		nbr *= -1
	}
	return nbr
}*/

//Fouth Variant
/*func main() {
	args := os.Args[1:]
	if len(os.Args) != 4 {
		return
	}
	a := atoi(args[0])
	b := atoi(args[2])
	if a == -999 || b == -999 {
		return
	}
	maxint := int(^uint(0) >> 1)
	minint := -maxint - 1
	if (a > 0 && b > 0 && (a+b) < 0) || a > maxint || b > maxint || a < minint || b < minint {
		return
	}
	switch args[1] {
	case "+":
		fmt.Printf("%d\n", a+b)
	case "-":
		fmt.Printf("%d\n", a-b)
	case "*":
		fmt.Printf("%d\n", a*b)
	case "/":
		if b != 0 {
			fmt.Printf("%d\n", a/b)
		} else {
			fmt.Printf("No division by 0\n")
		}
	case "%":
		if b != 0 {
			fmt.Printf("%d\n", a%b)
		} else {
			fmt.Printf("No modulo by 0\n")
		}
	default:
		return
	}
}

func atoi(s string) int {
	n := []rune(s)
	neg := false
	num := 0
	for i := 0; i < len(n); i++ {
		if i == 0 {
			if n[i] == '-' {
				neg = true
			}
		}
		if n[i] >= '0' && n[i] <= '9' {
			num = num*10 + int(n[i]) - 48
		} else if n[i] != '-' {
			return -999
		}
	}
	if neg == true {
		return -num
	}
	return num
} */

/* Another Variant of atoi

func atoi(s string) int {
	n := []rune(s)
	var flag bool
	num := 0
	for i := 0; i < len(n); i++ {
		if n[i] == '-' {
			flag = true
		} else if n[i] >= '0' && n[i] <= '9' {
			num = num*10 + int(n[i]) - 48
		} else {
			os.Exit(0)
		}
	}
	if flag {
		return num * -1
	}
	return num
} */

//Fifth Variant ("os" ONLY)
/*func main() {
 if len(os.Args) < 4 {
  return
 }
 value, operator, other := TrimAtoi(os.Args[1]), os.Args[2], TrimAtoi(os.Args[3])
 if value == 0 && os.Args[1] != "0"  other == 0 && os.Args[3] != "0"  len(os.Args) < 4 {
  return
 }
 if isOverflow(value) || isOverflow(other) {
  return
 }
 switch operator {
 case "+":
  printInt(value + other)
 case "-":
  printInt(value - other)
 case "*":
  printInt(value * other)
 case "/":
  if value == 0 || other == 0 {
   os.Stdout.WriteString("No division by 0")
  } else {
   printInt(value / other)
  }
 case "%":
  if value == 0 || other == 0 {
   os.Stdout.WriteString("No modulo by 0")
  } else {
   printInt(value % other)
  }
 default:
  return
 }
 os.Stdout.WriteString("\n")
}

func isOverflow(value int) bool {
 return value >= 1<<31 || value <= -(1<<31)
}

func printInt(nbr int) {
 os.Stdout.WriteString(ConvertNbr(nbr))
}

func ConvertNbr(nbr int) string {
 str := ""
 isNegative := nbr < 0
 for i, remainder := nbr, 0; i != 0; i /= 10 {
  remainder = i % 10
  if isNegative {
   remainder = -remainder
  }
  char := rune(remainder) + rune('0')
  str += string(char)
 }
 if isNegative {
  str += "-"
 } else if nbr == 0 {
  str = "0"
 }
 return Reverse(str)
}

func Reverse(s string) string {
 runes := []rune(s)
 for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
  runes[i], runes[j] = runes[j], runes[i]
 }
 return string(runes)
}

func TrimAtoi(s string) int {
 result := 0
 str := ""
 isNegative := false
 for index, char := range s {
  if index > 0 && s[index-1] == '-' && len(str) == 0 {
   isNegative = true
  }
  if isNumber(char) {
   str += string(char)
  }
 }
 str = Reverse(str)
 for index, char := range str {
  nb := int(char - '0')
  result += IterativePower(10, index) * nb
 }
 if isNegative {
  result = -result
 }
 return result
}

func isNumber(char rune) bool {
 return char >= '0' && char <= '9'
}

func IterativePower(nb int, power int) int {
 res := 1
 if power < 0 {
  return 0
 }
 for i := 1; i <= power; i++ {
  res = nb * res
 }
 return res
}*/
