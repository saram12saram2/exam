package piscine

func AtoiBase(s string, base string) int {
	set := make(map[rune]int)
	for i, j := range base {
		if _, k := set[j]; k || (j == '-' || j == '+') {
			return 0
		}
		set[j] = i
	}
	res := 0
	for i := 0; i < len(s); i++ {
		res = res*len(base) + set[rune(s[i])]
	}
	return res
}

//Second Variant
/*func AtoiBase(s string, base string) int {
	var res int
	l := len(base)
	if notDouble(base) || len(base) < 2 || base[0] == '-' || base[0] == '+' {
		return res
	}
	for i := range s {
		for j := range base {
			if s[i] == base[j] {
				res = res*l + j
			}
		}
	}
	return res
}
func notDouble(s string) bool {
	for i := range s {
		for k := 1; k < len(s); k++ {
			if s[i] == s[k] && i != k {
				return true
			}
		}
	}
	return false
}*/
