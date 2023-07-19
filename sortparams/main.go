package main

import ( "github.com/01-edu/z01"
		  "os" )
		  
func main() {
	arg := os.Args[1:]
	for i := 0; i < len(arg)-1; i++ {
		for j := i + 1; j < len(arg); j++ {
			if arg[i] > arg[j] {
				arg[i], arg[j] = arg[i], arg[j]
			}
		}
	}

	for _, args := range arg {
		for _, char := range args {
			z01.PrintRune(char)
		}
		z01.PrintRune('\n')
	}
}