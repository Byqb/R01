package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args[1:]
	for _, i := range arg {
		for _, w := range i {
			z01.PrintRune(w)
		}
		z01.PrintRune('\n')
	}
}
