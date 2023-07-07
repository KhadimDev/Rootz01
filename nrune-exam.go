package piscine 

func NRune(s string, n int) rune {
	c := []rune(s)
	for n <= len(s) {
		for i := 0; i < len(s); i++ {
			if i == (n - 1) {
				return rune(c[i])
			}
		}
	}
	return 0
}
