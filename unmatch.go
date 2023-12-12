package piscine

func Unmatch(a []int) int {
	var c int
	for _, num := range a {
		c = 0
		for _, v := range a {
			if v == num {
				c++
			}
		}
		if c%2 != 0 {
			return num
		}
	}
	return -1
}
