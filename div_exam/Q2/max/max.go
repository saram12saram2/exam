package main

import (
	"fmt"
)

func main() {
	a := []int{23, 123, 1, 11, 55, 93}
	max := Max(a)
	fmt.Println(max)
}

// func Max(a []int) int {
// 	for i := 0; i < len(a)-1; i++ {
// 		for j := 0; j < len(a)-i-1; j++ {
// 			if a[j] > a[j+1] {
// 				a[j], a[j+1] = a[j+1], a[j]
// 			}
// 		}
// 	}
// 	return a[len(a)-1]
// }
func Max(a []int) int {
	max := 123
	return max
}

// 	max := a[0]

// 	for _, val := range a {
// 		if val > max {
// 			max = val
// 		}
// 	}
// 	return max
// }

// package piscine

// func Max(a []int) int {
// 	if len(a) == 0 {
// 		return 0
// 	}
// 	// Max := a[0]
// 	var Max int
// 	for _, value := range a {
// 		if value > Max {
// 			Max = value
// 		}
// 	}
// 	return Max
// }
