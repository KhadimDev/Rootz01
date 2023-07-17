package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args[1:]
	for i := len(arg) - 1; i >= 0; i-- {
		value2 := arg[i]
		for _, value := range value2 {
			z01.PrintRune(value)
		}
		z01.PrintRune('\n')
	}
}
