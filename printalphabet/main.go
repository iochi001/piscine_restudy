package main

import "github.com/01-edu/z01"

func printAlphabetWhile() {
	letter := 'a'

	for letter <= 'z' {
		z01.PrintRune(letter)
		letter++
	}
	z01.PrintRune('\n')
}

func printAlphabetFor() {
	for letter := 'a'; letter <= 'z'; letter++ {
		z01.PrintRune(letter)
	}
	z01.PrintRune('\n')
}

func printAlphabetRange() {
	alphabet := "abcdefghijklmnopqrstuvwxyz"

	for _, letter := range alphabet {
		z01.PrintRune(letter)
	}
	z01.PrintRune('\n')
}

func main() {
	printAlphabetWhile()
	printAlphabetFor()
	printAlphabetRange()
}
