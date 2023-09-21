package piscine

func Split(s, sep string) []string {
	var res []string
	str := ""
	i := 0
	for len(s) > i {
		if s[i] == sep[0] {
			flag := s[i : len(sep)+i]
			if flag == sep {
				res = append(res, str)
				str = ""
				i += len(sep)
				continue
			}
		}
		str += string(s[i])
		i++
	}
	res = append(res, str)
	return res
}

// Second Variant
/*func Split(s, sep string) []string {
	var res []string
	word := ""
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(sep); j++ {
			if s[i] == sep[j] {
				if word != "" && i < len(s)-1 && j < len(sep)-1 && s[i+1] == sep[j+1] {
					res = append(res, word)
					word = ""
					i++
					break
				}
			}
			word += string(s[i])
			break
		}
	}
	if word != "" {
		res = append(res, word)
	}
	return res
}*/

// Third Variant
/*func Split(s, sep string) []string {
	sSep := s + sep
	count := 0
	for i := 0; i <= len(sSep)-len(sep); i++ {
		if sSep[i:i+len(sep)] == sep {
			count++
			i = i + len(sep)
		}
	}
	arr := make([]string, count)
	arrIdx, strIdx := 0, 0
	for i := 0; i <= len(sSep)-len(sep); i++ {
		if sSep[i:i+len(sep)] == sep {
			arr[arrIdx] = sSep[strIdx:i]
			strIdx = i + len(sep)
			arrIdx++
			i = i + len(sep)
		}
	}
	return arr
}*/
