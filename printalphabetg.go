package main

import "github.com/01-edu/z01"

func main() {
	for i := 'a'; i <= 'g'; i++ {
		z01.PrintRune(i)
	}
	z01.PrintRune('\n')
}