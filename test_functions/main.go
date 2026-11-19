package main

import (
	"piscine"

	"github.com/01-edu/z01"
)

func main() {
	piscine.IsNegative(1)
	piscine.IsNegative(0)
	piscine.IsNegative(-1)

	piscine.PrintComb()

	piscine.PrintComb2()

	piscine.PrintNbr(-123)
	piscine.PrintNbr(0)
	piscine.PrintNbr(123)
	z01.PrintRune('\n')

	piscine.PrintCombN(5)
}
