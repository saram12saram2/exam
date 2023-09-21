package main

import "github.com/01-edu/z01"

func main() {
	var oct byte = 108
	printBits(oct)
	printBits(swapBits(oct))
	printBits(reverseBits(oct))
}

func reverseBits(oct byte) byte {
	// getting an array of []int of binary representation of the given byte
	num := int(oct)
	arr := make([]int, 8)
	i := 7
	for num != 0 {
		arr[i] = num % 2
		num /= 2
		i--
	}
	var res int

	// reverse the values of the bits of a byte NOT THE BITS OF OUR BYTE
	bytes := []int{1, 2, 4, 8, 16, 32, 64, 128}

	// saving up the RES by multiplying the VALUES by the ONES AND ZEROS
	for i := range bytes {
		res += arr[i] * bytes[i]
	}

	return byte(res)
}

func swapBits(oct byte) byte {
	// getting an array of []int of binary representation of the given byte
	num := int(oct)
	arr := make([]int, 8)
	i := 7
	for num != 0 {
		arr[i] = num % 2
		num /= 2
		i--
	}
	var res int

	// swapping the values of the bits of a byte NOT THE BITS OF OUR BYTE
	bytes := []int{8, 4, 2, 1, 128, 64, 32, 16}

	// saving up the RES by multiplying the VALUES by the ONES AND ZEROS
	for i := range bytes {
		res += arr[i] * bytes[i]
	}
	return byte(res)
}

func printBits(oct byte) {
	num := int(oct)
	arr := make([]int, 8)
	i := 7
	for num != 0 {
		arr[i] = num % 2
		num /= 2
		i--
	}

	for i := range arr {
		z01.PrintRune(rune(arr[i] + '0'))
	}
	z01.PrintRune('\n')
}
