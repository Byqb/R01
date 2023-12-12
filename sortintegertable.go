package piscine

func SortIntegerTable(table []int) {
	n := len(table)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if table[i] < table[j] {
				s := table[i]
				table[i] = table[j]
				table[j] = s
			}
		}
	}
}
