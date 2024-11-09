
# def productExceptSelf(value: list[int]) -> list[int]:
# 	result = []
# 	check = [1] * (len(value))
# 	print(check)
# 	for i in value:
# 		exp = value[:]
# 		exp.remove(i)
# 		result.append(prod(exp))	
# 	return result

# def prod(val: list[int]) -> int:
# 	tot = 1
# 	for i in val:
# 		tot *= i
# 	return tot

# USING PREFIX AND SUFFIX INTUITION
def productExceptSelf(nums: list[int]) -> list[int]:
	res = [1] * (len(nums))

	prefix = 1
	for i in range(len(nums)):
		res[i] = prefix
		prefix *= nums[i]
		
	postfix = 1
	for i in range(len(nums) - 1, -1, -1):
		res[i] *= postfix
		postfix *= nums[i]
	return res

print(productExceptSelf([1,2,4,6]))