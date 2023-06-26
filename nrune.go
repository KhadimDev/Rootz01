package piscine

func NRune(s string, n int) rune {
	table := []rune(s)
	if n <= len(s) {
		for i := 0; i < len(s); i++ {
			if i == (n - 1) {
				return rune(table[i])
			}
		}
	}
	return 0
}
