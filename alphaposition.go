package piscine 

func AlphaPosition(c rune) int {
	var p rune
	if c >= 'A' && c <= 'Z' {
		p = c - 'A' + 1
		return int(p)
	} else if c >= 'a' && c <= 'z' {
		p = c - 'a' + 1
		return int(p)
	} else {
		return -1
	}
}
