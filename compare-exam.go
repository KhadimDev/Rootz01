package piscine 

func Compare(a, b string) int {
	var c int 
	if a > b {
		 return 1
	} else if a < b {
		return -1
	} else if a ==  b {
		return 0
	}
	return c
}
