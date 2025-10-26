package main

import (
	"fmt"
	"math/big"
)

// Function specifying returns for n = < 0 or 1 or 2.
func tribonacci(n int) []*big.Int {
	if n <= 0 {
		return []*big.Int{}
	}
	if n == 1 {
		return []*big.Int{big.NewInt(0)}
	}
	if n == 2 {
		return []*big.Int{big.NewInt(0), big.NewInt(1)}
	}
	// Start seq >= 3 with the first three terms

	fib_seq := []*big.Int{big.NewInt(1), big.NewInt(1), big.NewInt(1)}
	for i := 3; i < n; i++ {

		// Calculate the next terms and return.
		next_num := new(big.Int).Add(fib_seq[i-1], fib_seq[i-2])
		next_num.Add(next_num, fib_seq[i-3])
		fib_seq = append(fib_seq, next_num)
	}
	return fib_seq
}

func main() {
	n := 20 //Enter Test Number Here!
	fib_seq := tribonacci(n)

	//Ensuring the return for <=0 matches Python.
	if len(fib_seq) <= 0 {
		fmt.Print("[]")
	} else {
		//Ensureing the print format matches Python.
		fmt.Print("[")
		for x, next_num := range fib_seq {
			if x > 0 {
				fmt.Print(", ")
			}
			fmt.Print(next_num)
		}
		fmt.Print("]")
	}
}