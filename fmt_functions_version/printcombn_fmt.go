package main

import "fmt"

func PrintCombNFmt(n int) {
	if n <= 0 || n >= 10 {
		return
	}
	comb := make([]rune, n)
	printCombHelperFmt(0, '0', n, comb)
	fmt.Println()
}

func printCombHelperFmt(pos int, start rune, n int, comb []rune) {
	if pos == n {
		fmt.Printf("%s", string(comb))
		if comb[0] != rune('9'-n+1) {
			fmt.Print(", ")
		}
		return
	}

	for d := start; d <= '9'; d++ {
		comb[pos] = d
		printCombHelperFmt(pos+1, d+1, n, comb)
	}
}
