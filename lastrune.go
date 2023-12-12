package piscine

func LastRune(s string) rune {
	a := []rune(s)
	length := len(a) - 1
	return a[length]
}
