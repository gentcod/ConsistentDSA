package averagewaitingtime

import (
	"fmt"
)

// Leetcode Problem: https://leetcode.com/problems/average-waiting-time/
func Solve(cs [][]int) {
	fmt.Println(averageWaitingTime(cs))
}

func averageWaitingTime(cs [][]int) float64 {
	cur := cs[0][0] + cs[0][1]
	twt := cur - cs[0][0]

	for i := range cs {
		if i == 0 {
			continue
		}
		if cur < cs[i][0] {
			cur = cs[i][0] + cs[i][1]
		} else {
			cur += cs[i][1]
		}
		twt += (cur - cs[i][0])
	}

	return (float64(twt) / float64(len(cs)))
}
