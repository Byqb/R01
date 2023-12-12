package piscine

func DescendAppendRange(max, min int) []int {
	answer := []int{}
	if max <= min {
		return []int{}
	} else {
		for i := max; i >= min+1; i-- {
			answer = append(answer, i)
		}
	}
	return answer
}
