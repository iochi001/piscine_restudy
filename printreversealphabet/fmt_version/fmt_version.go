package main

import "fmt"

func printAlphabetWhile1() {
	letter := 'z'

	for letter >= 'a' {
		fmt.Printf("%c", letter)
		letter--
	}
	fmt.Println()
}

func printAlphabetFor1() {
	for letter := 'z'; letter >= 'a'; letter-- {
		fmt.Printf("%c", letter)
	}
	fmt.Println()
}

func printAlphabetRange1() {
	alphabet := "zyxwvutsrqponmlkjihgfedcba"

	for _, letter := range alphabet {
		fmt.Printf("%c", letter)
	}
	fmt.Println()
}

func main() {
	printAlphabetWhile1()
	printAlphabetFor1()
	printAlphabetRange1()
}
