package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	var ArrayOfStrings []string
	EachStringAlone := ""
	str = str + " "
	for _, x := range str {
		if x == ' ' {
			ArrayOfStrings = append(ArrayOfStrings, EachStringAlone)
			EachStringAlone = ""
		} else {
			EachStringAlone += string(x)
		}
	}
	StringAndItsValue := make(map[string]int)
	for _, EachWord := range ArrayOfStrings {
		_, TrueOrFalse := StringAndItsValue[EachWord]
		if TrueOrFalse {
			StringAndItsValue[EachWord]++
		} else {
			StringAndItsValue[EachWord] = 1
		}
	}
	return StringAndItsValue
}
