package doublemodularexp

import (
	"fmt"
	// "math/big"
)

func Solve(variables [][]int, target int) {
	fmt.Println(getGoodIndices(variables, target))
}

func getGoodIndices(variables [][]int, target int) (result []int) {
	for i, v := range variables {
		a, b, c, m := v[0], v[1], v[2], v[3]

		if (powmod(powmod(a, b, 10), c, m)) == target {
			result = append(result, i)
		}

		// this has issues because Go has an upper bound of 32 or 64 bits for integers
		// if pow((pow(a,b) % 10), c) % m == target {
		// 	result = append(result, i)
		// }
	}
	return
}

func powmod(base, exp, mod int) int {
	if exp == 0 {
		return 1
	}
	half := powmod(base, exp/2, mod)
	half = (half * half) % mod
	if exp%2 != 0 {
		half = (half * base) % mod
	}
	return half
}

// this provides a solution to the upper bound of integers
// func powmod(base, exp, mod int64) *big.Int {
// 	result := new(big.Int)

// 	// Perform (base^exp) % mod using math/big's Exp function
// 	result.Exp(big.NewInt(base), big.NewInt(exp), big.NewInt(mod))

// 	return result
// }

// this returns the power of a base based on an exponential.. math.Pow() implementation for int data types
// func pow(base, exp int) int {
// 	result := 1
// 	for exp > 0 {
// 		result = result * base
// 		exp--
// 	}

// 	return result
// }