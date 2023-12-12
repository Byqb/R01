package piscine

func ToLower(s string) string {
	a := []rune(s)
	for i := range a {
		if a[i] >= 'A' && a[i] <= 'Z' {
			a[i] = a[i] + ' '
		}
	}
	return string(a)
}
