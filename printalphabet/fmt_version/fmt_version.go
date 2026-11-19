package main

import "fmt"

func printAlphabetWhile1() {
	letter := 'a'

	for letter <= 'z' {
		fmt.Printf("%c", letter)
		letter++
	}
	fmt.Println()
}

func printAlphabetFor1() {
	for letter := 'a'; letter <= 'z'; letter++ {
		fmt.Printf("%c", letter)
	}
	fmt.Println()
}

func printAlphabetRange1() {
	alphabet := "abcdefghijklmnopqrstuvwxyz"

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
