package main

import "fmt"

func IsNegativeFmt(nb int) {
	if nb < 0 {
		fmt.Printf("%c\n", 'T')
	} else {
		fmt.Printf("%c\n", 'F')
	}
}
