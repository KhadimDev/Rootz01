package main

func AlphaCount(str string) int {
	c := 0
	count := []rune(str)
	for i := 0; i < len(str); i++ {
		if count[i] >= 'A' && count[i] <= 'Z' || count[i] >= 'a' && count[i] <= 'z' {
			c++
		}
	}
	return c
}
