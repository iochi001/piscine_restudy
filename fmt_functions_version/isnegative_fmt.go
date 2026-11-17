package main

import "fmt"

func IsNegative(nb int) {
	if nb < 0 {
		fmt.Printf("%c\n", 'T')
	} else {
		fmt.Printf("%c\n", 'F')
	}
}

func main() {
	IsNegative(1)
	IsNegative(0)
	IsNegative(-1)
}
