package main

import ("os"

	   "github.com/01-edu/z01"
)

func main() {
	arg := os.Args[1:]
	for _, res := range arg {
		for i, j := range res {
			if i >= 0 {
				z01.PrintRune(j)
			}
		}
		z01.PrintRune('\n')
	}
}
