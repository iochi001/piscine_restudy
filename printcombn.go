package piscine

import "github.com/01-edu/z01"

func PrintCombN(n int) {
	if n <= 0 || n >= 10 {

		z01.PrintRune('\n')
		return
	}
	comb := make([]rune, n)
	printCombHelper(0, '0', n, comb)
	z01.PrintRune('\n')
}

func printCombHelper(pos int, start rune, n int, comb []rune) {
	if pos == n {
		for i := 0; i < n; i++ {
			z01.PrintRune(comb[i])
		}

		if comb[0] != rune('9'-n+1) {
			z01.PrintRune(',')
			z01.PrintRune(' ')
		}
		return
	}

	// Optimized loop bound
	maxDigit := '9' - rune(n-1-pos)
	for d := start; d <= maxDigit; d++ {
		comb[pos] = d
		printCombHelper(pos+1, d+1, n, comb)
	}
}
