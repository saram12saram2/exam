package piscine

func SwapBits(octet byte) byte {
	return octet<<4 + octet>>4
}

// Tester
/*func main() {
	var b byte = 65
	result := int(SwapBits(b))
	// Print Byte
	slice := make([]int, 8)
	i := len(slice) - 1
	for result != 0 {
		slice[i] = result % 2
		result /= 2
		i--
	}
	fmt.Println(slice)
}*/

// Second Variant
/*func SwapBits(fool byte) byte {
	temp := int(fool)
	var result int
	slice := make([]int, 8)
	// from int to byte
	i := len(slice) - 1
	for temp != 0 {
		slice[i] = temp % 2
		temp /= 2
		i--
	}
	// from byte to int
	value := []int{8, 4, 2, 1, 128, 64, 32, 16}
	for i := range slice {
		if slice[i] == 1 {
			result += value[i]
		}
	}
	return byte(result)
}*/
