package integertoroman

import "fmt"

func Solve(n int) {
	fmt.Println(integerToRoman(n))
}

func integerToRoman(n int) (r string) {
	keys := []int{1000,500,100,50,10,5,1}
	literals := map[int]string{
		1000: "M",
		500: "D",
		100: "C",
		50: "L",
		10: "X",
		5: "V",
		1 : "I",
	}

	for k,i := range(keys) {
		if n % i != n {
			div := (n - (n % i)) / i
			if div == 4 {
				if len(r) > 0 && literals[keys[k-1]] == r[len(r)-1:] {
					r = r[:len(r)-1]
					r += literals[i] + literals[keys[k-2]]
				} else {
					r += literals[i] + literals[keys[k-1]]
				}
			} else {
				r += evalLiteral(div, literals[i])
			}
		}
		n = n % i
	}

	return
}

func evalLiteral (div int, l string) (r string) {
	for i := div; i > 0; i-- {
		r += l
	}

	return
}
