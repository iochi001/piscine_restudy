package main

import (
	"fmt"
	"math"
)

func main() {
	IsNegativeFmt(1)
	IsNegativeFmt(0)
	IsNegativeFmt(-1)

	PrintCombFmt()

	PrintComb2Fmt()

	PrintNbrFmt(-123)
	fmt.Println()
	PrintNbrFmt(0)
	fmt.Println()
	PrintNbrFmt(123)
	fmt.Println()
	PrintNbrFmt(math.MinInt)
	fmt.Println()

	PrintCombNFmt(1)
	PrintCombNFmt(3)
	PrintCombNFmt(9)
}
