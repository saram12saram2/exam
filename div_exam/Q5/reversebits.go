package piscine

// Reversebits(00100110) -> 01100100
func ReverseBits(oct byte) byte {
	var res byte
	for i := 8; i > 0; i-- {
		res = (res << 1) | (oct & 1)
		oct >>= 1
	}
	return res
}

// Tester
/*func main() {
	var b byte = 38
	fmt.Println(b)
	result := int(ReverseBits(b))
	// Print Byte
	slice := make([]int, 8)
	i := len(slice) - 1
	for result != 0 {
		slice[i] = result % 2
		result /= 2
		i--
	}
	fmt.Println(slice)
	//for _, w := range slice {
	//	z01.PrintRune(rune(w) + 48)
	//}
	//z01.PrintRune('\n')
}*/

// Second Variant
/*func ReverseBits(oct byte) byte {
	// converting byte to an int for convenience
	var given int = int(oct)
	// making a array with size of 8
	slice := make([]int, 8)
	// from int to an array of int
	i := len(slice) - 1
	for given != 0 {
		slice[i] = given % 2
		given /= 2
		i--
	}
	// converting the byte array to one int in the reversed order
	var res int
	value := []int{128, 64, 32, 16, 8, 4, 2, 1}
	for i := range slice {
		if slice[i] == 1 {
			res += value[i]
		}
	}
	// if slice[0] == 1 {
	// 	res += 128
	// }
	// if slice[1] == 1 {
	// 	res += 64
	// }
	// if slice[2] == 1 {
	// 	res += 32
	// }
	// if slice[3] == 1 {
	// 	res += 16
	// }
	// if slice[4] == 1 {
	// 	res += 8
	// }
	// if slice[5] == 1 {
	// 	res += 4
	// }
	// if slice[6] == 1 {
	// 	res += 2
	// }
	// if slice[7] == 1 {
	// 	res += 1
	// }
	// converting int to byte and returning the result
	return byte(res)
}*/
