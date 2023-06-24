package piscine 

func Sqrt(nb int) int {
	if nb < 0 {
		return 0 
	}
	root := 0
	for i := 0; i*i <= nb; i++ {
		root = i
	}
	if root*root == nb {
		return root
	}
	return 0 
}
