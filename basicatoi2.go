package piscine

func BasicAtoi2(s string) int {
	r := 0
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			r = r*10 + int(s[i]-'0')
		} else {
			return 0
		}
	}
	return r 
}
