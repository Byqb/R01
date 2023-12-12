package piscine

func SplitWhiteSpaces(s string) []string {
	var str []string
	word := ""
	for _, l := range s {
		if l == ' ' {
			if len(word) > 0 {
				str = append(str, word)
				word = ""
			}
		} else {
			word += string(l)
		}
	}
	if len(word) > 0 {
		str = append(str, word)
	}
	return str
}
