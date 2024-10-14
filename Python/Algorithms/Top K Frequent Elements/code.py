def topKFrequent(nums: list[int], k: int):
   hashMap = {}
   maxoccur = 0
   result = []

   for i in range(len(nums)):
      if nums[i] in hashMap:
         hashMap[nums[i]] = hashMap.get(nums[i]) + 1
      else: hashMap[nums[i]] = 1

      if hashMap[nums[i]] > maxoccur:
         maxoccur = hashMap[nums[i]]

   freq: list[int] = [[] for _ in range(maxoccur+1)]
   
   for i in hashMap:
      freq[hashMap[i]].append(i)

   for i in range(k):
      for j in freq[maxoccur]:
         if len(result) == k:
            return result
         result.append(j)

      maxoccur -= 1

   return result

nums = [1,1,1,2,2,3,3,3]
print('Answer: ',topKFrequent(nums,2))