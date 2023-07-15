package main 

import ("os"

	   "github.com/01-edu/z01"
)
func main() {
	m := []rune(os.Args[0])
	for i := 2; i < len(m); i++ {
		z01.PrintRune(m[i])
	}
	z01.PrintRune('\n')
}