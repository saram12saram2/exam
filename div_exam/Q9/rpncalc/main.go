package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Error")
		return
	}
	var vals []int
	op := strings.Split(os.Args[1], " ")
	op = Del(op)
	for _, v := range op {
		val, err := strconv.Atoi(v)

		if err == nil {
			vals = append(vals, val)
			continue
		}

		n := len(vals)
		if n < 2 {
			fmt.Println("Error")
			return
		}

		switch v {
		case "+":
			vals[n-2] += vals[n-1]
			vals = vals[:n-1]
		case "-":
			vals[n-2] -= vals[n-1]
			vals = vals[:n-1]
		case "*":
			vals[n-2] *= vals[n-1]
			vals = vals[:n-1]
		case "/":
			vals[n-2] /= vals[n-1]
			vals = vals[:n-1]
		case "%":
			vals[n-2] %= vals[n-1]
			vals = vals[:n-1]
		default:
			fmt.Println("Error")
			return
		}
	}
	if len(vals) == 1 {
		fmt.Println(vals[0])
	} else {
		fmt.Println("Error")
	}
}

func Del(a []string) (res []string) {
	for _, v := range a {
		if v != "" {
			res = append(res, v)
		}
	}
	return
}
