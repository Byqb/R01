package main

import (
	"fmt"
	"os"
	"strconv"
)

func is0pr(s string, array []string) bool {
	for _, ops := range array {
		if ops == s {
			return true
		}
	}
	return false
}

func main() {
	oper := []string{"+", "-", "*", "/", "%"}
	arg := os.Args[1:]

	if len(arg) != 3 {
		fmt.Print()
	} else {
		if is0pr(arg[1], oper) {
			first, err1 := strconv.Atoi(arg[0])
			second, err2 := strconv.Atoi(arg[2])
			if first > 10000 || second > 10000 {
				return
			}
			if err1 == nil && err2 == nil {
				switch arg[1] {
				case "+":
					fmt.Println(first + second)
				case "-":
					fmt.Println(first - second)
				case "*":
					fmt.Println(first * second)
				case "/":
					if second == 0 {
						fmt.Println("No division by 0")
					} else {
						fmt.Println(first / second)
					}
				case "%":
					if second == 0 {
						fmt.Println("No modulo by 0")
					} else {
						fmt.Println(first % second)
					}
				}
			} else {
				fmt.Print()
			}
		} else {
			fmt.Print()
		}
	}
}
