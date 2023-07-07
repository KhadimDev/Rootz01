package piscine 

func CountAlpha(s string) int {
	compte := 0
	c := []rune(s)

	for i := 0; i < len(s); i++ {
		if c[i] >= 'a' && c[i] <= 'z' || c[i] >= 'A' && c[i] <= 'Z' {
			compte++
		}
	}
	return compte 
}
