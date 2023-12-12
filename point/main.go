package main

import "github.com/01-edu/z01"

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func PrintInt(x rune) {
	z01.PrintRune(x)
}

func main() {
	points := &point{}
	setPoint(points)
	d := "x = "
	f := ", y = "
	for _, cx := range d {
		z01.PrintRune(cx)
	}
	PrintInt(52) // 4
	PrintInt(50) // 2
	for _, cy := range f {
		z01.PrintRune(cy)
	}
	PrintInt(50) // 2
	PrintInt(49) // 1
	z01.PrintRune('\n')
}
