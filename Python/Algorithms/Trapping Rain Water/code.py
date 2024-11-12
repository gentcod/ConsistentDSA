# PROBLEM lEETCODE - https://leetcode.com/problems/trapping-rain-water/description/

def trappingwater(height: list[int]) -> int:
	l, r = 0, len(height) - 1
	lMax, rMax = height[l], height[r]
	res = 0

	while l < r:
		if lMax < rMax:
			l += 1
			lMax = max(lMax, height[l])
			res += lMax - height[l]
		else:
			r -= 1
			rMax = max(rMax, height[r])
			res += rMax - height[r]

	return res

print(trappingwater([4,2,0,3,2,5]))