package main

import "fmt"

func printDigitsWhile1() {
	number := '0'

	for number <= '9' {
		fmt.Printf("%c", number)
		number++
	}
	fmt.Printf("\n")
}

func printDigitsFor1() {
	for number := '0'; number <= '9'; number++ {
		fmt.Printf("%c", number)
	}
	fmt.Printf("\n")
}

func printDigitsRange1() {
	number := "0123456789"

	for _, number := range number {
		fmt.Printf("%c", number)
	}
	fmt.Printf("\n")
}

func main() {
	printDigitsWhile1()
	printDigitsFor1()
	printDigitsRange1()
}
