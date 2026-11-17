package main

import "github.com/01-edu/z01"

func printDigitsWhile() {
number := '0'

  for number <= '9' {
    z01.PrintRune(number)
    number++
  }
z01.PrintRune('\n')
}

func printDigitsFor() {
 for number := '0'; number <= '9'; number++ {
  z01.PrintRune(number)
 }
 z01.PrintRune('\n')
}

func printDigitsRange() {
  number := "0123456789"

  for _, number := range number {
    z01.PrintRune(number)
  }
  z01.PrintRune('\n')
}

func main() {
	printDigitsWhile()
	printDigitsFor()
  printDigitsRange()
}