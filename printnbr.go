package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	// 2. Handle negative numbers
	if n < 0 {
		z01.PrintRune('-')
		n = -n // Flip it to positive
	}

	// 3. The Recursive Logic
	if n >= 10 {
		// If it has multiple digits, chop off the last digit and call ourselves
		PrintNbr(n / 10)
	}

	// 4. Print the current digit (Base case also ends up here)
	z01.PrintRune(rune(n%10) + '0')
}
