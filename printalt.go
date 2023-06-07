package main 

import "github.com/01-edu/z01"

func main() {
	for i, j := 'a', 'z'; j >= 'n'; i, j = i + 1, j - 1 {
		z01.PrintRune(i)
		z01.PrintRune(j)
	}
	z01.PrintRune('\n')
}