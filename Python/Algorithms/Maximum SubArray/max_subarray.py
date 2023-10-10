def maxSubArray(nums) -> int:
   maxSub = nums[0]
   curSum = 0

   for i in nums:
      if curSum < 0:
         curSum = 0
      
      curSum += i
      maxSub = max(maxSub, curSum)
      
   return maxSub
   
# print(maxSubArray([-1, 2, - 3, -4, 4, 2, -1, 1, 5, -12]))
# print(maxSubArray([-2,1,-3,4,-1,2,1,-5,4]))