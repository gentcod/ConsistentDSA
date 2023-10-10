def twoSum(arr, target):
   hashMap = {}

   for i, n in enumerate(arr):
      diff = target - n
      print(hashMap)
      if diff in hashMap:
         return [hashMap[diff], i]
      
      hashMap[n] = i

# print(twoSum([2, 3, 4, 5], 7))