package guesshigherorlower

import (
	"fmt"
)

// Leetcode Problem: https://leetcode.com/problems/guess-number-higher-or-lower/
func Solve(n,pick int) {
	fmt.Println(guessNumber(n,pick))
}

func guessNumber(n,pick int) int {
	return binarySearch(0,n,pick)
}

func binarySearch(l,r,pick int) int {
  if l > r {
	  return -1
  }

	mid := (l + r) / 2
	if guess(mid, pick) == -1 {
		 return binarySearch(l,mid-1,pick)
	} 
	if guess(mid, pick) == 1 {
		 return binarySearch(mid+1,r,pick)
	}
	if guess(mid, pick) == 0 {
		 return mid
	}

  return mid
}

func guess(num,pick int) int {
	if num > pick {
		return -1
	}
	if num < pick {
		return 1
	}

	return 0
}