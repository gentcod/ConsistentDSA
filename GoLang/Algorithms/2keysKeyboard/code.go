package twokeyskeyboard

import "fmt"

func Solve(n int) {
	fmt.Println(minSteps(n))
}

func minSteps(n int) int {
	if n == 1 {
		return 0;
	}
	// FIND THE HIGHIEST DIVISOR, RECURSIVE ADD n / DIVISOR
	for i := n >> 1; i > 0; i-- {
		if (n % i) == 0 {
			fmt.Println("nval: ", n,"div: ",i, "calc: ", n / i)
			return n / i + minSteps(i)
		}
	}
	return -1
}

// INITIAL
// func minSteps(n int) int {
// 	if n == 1 {
// 		return 0
// 	}

// 	div := n
// 	target := 1
// 	clip := 1
// 	copy := 1
// 	copied := true
// 	paste := 0
// 	for i := 0; target != n; i++ {
// 	// for i := 0; i < 6; i++ {
// 		if div % target == 0 && target > 1 && !copied{
// 			copy++
// 			div = div >> 1
// 			clip = target
// 			copied = true
// 		} else {
// 			target += clip
// 			paste++
// 			copied = false
// 		}

// 		fmt.Println("div:", div)
// 		fmt.Println("target:", target)
// 		fmt.Println("clip:", clip)
// 		fmt.Println("paste:", paste)
// 		fmt.Println("copy:", copy)
// 		fmt.Println("----------------")
// 	}

// 	return copy + paste
// }
