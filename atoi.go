package piscine

func Atoi(s string) int {
	num := 0
	multiplicateur := 1
	i := 0

	if len(s) < 0 {
		return 0
	}

	if s[0] == '-' {
		multiplicateur = -1
		i++
	}
	if s[0] == '+' {
		multiplicateur = 1
		i++
	}

	for i < len(s) {
		num = num*10 + int(s[i]-'0')
		if (int(s[i]-'0')) > 9 || (int(s[i]-'0')) < 0 {
			return 0
		}
		i++
	}
	return num * multiplicateur
}
