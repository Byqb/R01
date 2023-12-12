package piscine

func Capitalize(s string) string {
	c := true
	sb := []rune(s)
	for i, b := range sb {
		if (b >= 'a' && b <= 'z') && c {
			sb[i] = sb[i] - 32
		} else if (b >= 'A' && b <= 'Z') && !c {
			sb[i] = sb[i] + 32
		}
		if (b < 'A' || b > 'Z') && (b < 'a' || b > 'z') && (b < '0' || b > '9') {
			c = true
		} else {
			c = false
		}
	}
	return string(sb)
}
