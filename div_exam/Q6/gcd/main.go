package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) == 3 {
		n1, _ := strconv.Atoi(os.Args[1])
		n2, _ := strconv.Atoi(os.Args[2])

		if n1 < 0 || n2 < 0 {
			return
		}
		if n1 < n2 {
			n1, n2 = n2, n1
		}
		for n2 != 0 { // n2 > 0
			n1, n2 = n2, n1%n2
		}
		fmt.Println(n1)
	}
}

// Second Variant
/*func main() {
	if len(os.Args) != 3 {
		return
	}
	num1, err := strconv.Atoi(os.Args[1])
	num2, err := strconv.Atoi(os.Args[2])
	if err != nil {
		return
	}
	gcd(num1, num2)
}

func gcd(n1, n2 int) {
	var num int
	for i := 1; i <= n1 && i <= n2; i++ {
		if n1%i == 0 && n2%i == 0 {
			num = i
		}
	}
	fmt.Println(num)
}*/
