package piscine 

func Compare(a, b string) int {
	var num int
	if a == b {
		num = 0
	} else if a > b {
		num = 1
	} else if a < b {
		num = -1
	}
	return num 
}
