package piscine

func ToUpper(s string) string {
	a := []rune(s)
	for i := range a {
		if a[i] >= 'a' && a[i] <= 'z' {
			a[i] = a[i] - ' '
		}
	}
	return string(a)
}
