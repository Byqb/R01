package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	a := os.Args[0]
	for i, p := range a {
		if i > 1 {
			z01.PrintRune(p)
		}
	}
	z01.PrintRune('\n')
}
