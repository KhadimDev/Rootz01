package main

import "fmt"

func BasicJoin(elems []string) string {
	var c string
	for _, r := range elems {
		c += r
	}
	return c
}

func main() {
	elems := []string{"Hello!", " How", " are", " you?"}
	fmt.Println(BasicJoin(elems))
}