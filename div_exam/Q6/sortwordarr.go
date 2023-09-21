package piscine

func SortWordArr(a []string) {
	for i := 0; i < len(a); i++ {
		for j := i; j < len(a); j++ {
			if a[i] > a[j] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
}

//Second Variant
/*func SortWordArr(a []string) {
    for i := 0; i < len(a)-1; i++ {
        for j := i; j < len(a)-i-1; j++ {
            if a[j] > a[j+1] {
                a[j], a[j+1] = a[j+1], a[j]
            }
        }
    }
}*/
