package saga

import "fmt"

func ReverseTransaction(ref string) {
	fmt.Println("Reversing transaction:", ref)

	// 1. Credit wallet
	// 2. Update transaction status to REVERSED
}
