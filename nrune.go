package piscine

func NRune(s string, n int) rune {
	a := []rune(s)
	if n-1 >= 0 && n-1 < len(a) {
		return a[n-1]
	} else {
		return 0
	}
}
