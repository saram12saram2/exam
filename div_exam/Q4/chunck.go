package piscine

import "fmt"

func Chunk(slice []int, size int) {
	bs := [][]int{}
	if len(slice) == 0 {
		fmt.Println() // fmt.Println(bs)
		return
	}
	if size <= 0 {
		fmt.Println()
		return
	}

	for len(slice) > size {
		bs = append(bs, slice[:size])
		slice = slice[size:]
	}

	bs = append(bs, slice)
	fmt.Println(bs)
}

//Second Variant
/*func Chunk(slice []int, size int) {
    bs := [][]int{}

    for len(slice) > size {
        if size != 0 {
            bs = append(bs, slice[:size])
            slice = slice[size:]
        } else {
            break
        }
    }

    if len(slice) == 0 {
        fmt.Println(bs)
        return
    }

    if size <= 0 {
        fmt.Println()
        return
    }

    bs = append(bs, slice)
    fmt.Println(bs)
}*/
