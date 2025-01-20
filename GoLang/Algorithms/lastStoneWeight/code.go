package laststoneweight

import (
	"fmt"
	"sort"
)

// Leetcode Problem: https://leetcode.com/problems/last-stone-weight/description/
func Solve(stones []int) {
	fmt.Println(lastStoneWeight(stones))
}

func lastStoneWeight(stones []int) int {
	if len(stones) == 1 {
		return stones[0]
	}

	sort.Ints(stones)
	for len(stones) > 0 {
        if len(stones) == 1 {
			return stones[0]
		}
		l, r := len(stones)-1, len(stones)-2
		if stones[l] == stones[r] {
			stones = stones[:r]
		} else {
			stones[r] = stones[l] - stones[r]
			stones = stones[:l]
	        sort.Ints(stones)
		}
	}

	return 0
}