package main 

import "github.com/01-edu/z01"

func printrevcomb() {
	for i := '9'; i >= '2'; i-- {
		for j := '8'; j >= '1'; j-- {
			for k := '7'; k >= '0'; k-- {
				if i > j && j > k {
					z01.PrintRune(i)
					z01.PrintRune(j)
					z01.PrintRune(k)
					if i > 2 {
						z01.PrintRune(',')
						z01.PrintRune(' ')
					}
				}
			} 
		}
	}
	z01.PrintRune('\n')
}

func main() {
	printrevcomb()
}