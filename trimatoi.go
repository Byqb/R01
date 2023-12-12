package piscine

func TrimAtoi(s string) int {
	T := 0
	G := 1
	for _, c := range s {
		if c >= '0' && c <= '9' {

			b := int(c - 48)
			T = T*10 + b
		} else if c == '-' && T == 0 {
			G = -1
		}
	}
	return T * G
}
