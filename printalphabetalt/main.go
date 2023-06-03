package main 

import "github.com/01-edu/z01"

func main() {
	for i, j := 'a', 'B'; j <= 'Z'; i, j = i + 2, j + 2 {
		z01.PrintRune(i)
		z01.PrintRune(j)
	}
	z01.PrintRune('\n')
	for i, j := 'z', 'Y'; j >= 'A'; i, j = i - 2, j - 2 {
		z01.PrintRune(i)
		z01.PrintRune(j)
	}
	z01.PrintRune('\n')
}