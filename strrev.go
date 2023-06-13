package piscine 

func StrRev(s string) string {
	p := []rune(s)
	for i, j := 0, len(s) - 1; i < j; i, j = i + 1, j - 1 {
		p[i], p[j] = p[j], p[i]
	}
	return string(p)
}
