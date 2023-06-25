package piscine 

func LastRune(s string) rune {
	table := []rune(s)
	return table[len(s)-1]
}
