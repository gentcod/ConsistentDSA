def twoSum(arr, target):
   hashMap = {}

   for i, n in enumerate(arr):
      diff = target - n
      if diff in hashMap:
         return [hashMap[diff], i]
      
      hashMap[n] = i

print(twoSum([4, 5, 3, 2], 7))