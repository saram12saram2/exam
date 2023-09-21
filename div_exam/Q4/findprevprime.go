package piscine

func FindPrevPrime(nb int) int {
	if nb <= 1 {
		return 0
	}
	for nb > 1 {
		if isprime(nb) {
			break
		}
		nb--
	}
	return nb
}

func isprime(nb int) bool {
	for i := 2; i <= nb/2; i++ {
		if nb%i == 0 {
			return false
		}
	}
	return true
}



//Second Variant
/*func FindPrevPrime(nb int) int {
    for i := nb; i > 2; i-- {
        if IsPrime(i) {
            return i
        }
    }
    return 0
}

func IsPrime(n int) bool {
    for i := 2; i < n; i++ {
        if n%i == 0 {
            return false
        }
    }
    return true
}*/
