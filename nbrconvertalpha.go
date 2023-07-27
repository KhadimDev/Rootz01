package main

import (
	"os"
	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:] // Get command-line arguments excluding the program name

	upperCase := false
	if len(args) > 0 && args[0] == "--upper" {
		upperCase = true
		args = args[1:] // Remove the "--upper" flag from the arguments
	}

	for _, arg := range args {
		num := 0
		for _, c := range arg {
			if c < '0' || c > '9' {
				num = -1
				break
			}
			num = num*10 + int(c-'0')
		}

		if num < 1 || num > 26 {
			z01.PrintRune(' ')
			continue
		}

		letter := 'a' + rune(num-1)
		if upperCase {
			letter -= 32 // Convert to uppercase by subtracting 32 from the lowercase letter's ASCII value
		}

		z01.PrintRune(letter)
	}
	z01.PrintRune('\n') // Print a new line after all letters are printed
}