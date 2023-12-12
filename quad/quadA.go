package piscine

import "github.com/01-edu/z01"

func QuadA(x, y int) {
	if x <= 0 || y <= 0 { //we added an if condition stating that x should be less than or equal to 0 or y is less than or equal to 0 to make sure that the int is positive.
	} else {
		for a := 1; a <= y; a++ { //we defined a as 1 and we set the condition that a should be less than or equal to y, if that condition was met the a +1
			for b := 1; b <= x; b++ { //we defined b as 1 and we set the condition that a should be less than or equal to x, if that condition was met the b +1
				if a == y && b == x || a == 1 && b == 1 || a == y && b == 1 || a == 1 && b == x { // we set an if condition stating that if a is equal to y and b is equal to y
					z01.PrintRune('o') // print o
				} else if b == 1 && (a != y || a != 1) { //if b == 1 but a doesn not equal to y or 1
					z01.PrintRune('|') // print |
				} else if b == x && (a != 1 || a != y) { //if b == x but a doesn not equal to y or 1
					z01.PrintRune('|') // print |
				} else if a == 1 && (b != 1 || b != x) { //if a == 1 but b doesn not equal to x or 1
					z01.PrintRune('-') // print -
				} else if a == y && (b != 1 || b != x) { //if a == y but b doesn not equal to x or 1
					z01.PrintRune('-') // print -
				} else {
					z01.PrintRune(' ') // print nothing
				}
			}
			z01.PrintRune('\n') // new line
		}
	}
}
