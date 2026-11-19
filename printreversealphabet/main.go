package main

import "github.com/01-edu/z01"

func printReverseAlphabetWhile() {
	letter := 'z'

	for letter >= 'a' {
		z01.PrintRune(letter)
		letter--
	}
	z01.PrintRune('\n')
}

func printReverseAlphabetFor() {
	for letter := 'z'; letter >= 'a'; letter-- {
		z01.PrintRune(letter)
	}
	z01.PrintRune('\n')
}

func printReverseAlphabetRange() {
	alphabet := "zyxwvutsrqponmlkjihgfedcba"

	for _, letter := range alphabet {
		z01.PrintRune(letter)
	}
	z01.PrintRune('\n')
}

func main() {
	printReverseAlphabetWhile()
	printReverseAlphabetFor()
	printReverseAlphabetRange()
}
