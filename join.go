package piscine

func Join(strs []string, sep string) string {
	var c string
	for i, r := range strs {
		c += r 
		if i != len(strs)-1 {
			c += sep
		}
	}
	return c 
}
