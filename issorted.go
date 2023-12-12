package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	ac := true
	de := true
	for i := 0; i < len(a)-1; i++ {
		if f(a[i], a[i+1]) < 0 {
			de = false
		}
	}
	for i := 0; i < len(a)-1; i++ {
		if f(a[i], a[i+1]) > 0 {
			ac = false
		}
	}

	return ac || de
}
