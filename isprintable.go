package piscine 

func IsPrintable(s string) bool {
	c := []rune(s)
	for i := 0; i < len(s); i++ {
		if c[i] >= '\a' && c[i] <= '\r' {
			return false
		}
	}
	return true
}
