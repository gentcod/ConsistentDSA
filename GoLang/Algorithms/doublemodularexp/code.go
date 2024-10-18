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

// SAMPLE DATA
// variables := [][]int{{2,3,3,10},{3,3,3,1},{6,1,1,4}}
// variables := [][]int{
// 	{30,5,43,2},{15,50,35,41},{45,34,41,32},{14,37,33,13},{6,8,1,53},{37,1,12,52},{42,37,2,52},{9,2,15,3},{31,12,21,24},{52,24,6,12},
// 	{51,35,21,52},{30,18,10,2},{27,31,50,27},{29,25,26,32},{15,38,43,17},{22,12,16,43},{48,9,15,6},{41,26,22,21},{41,49,52,26},{53,38,9,33}
// }