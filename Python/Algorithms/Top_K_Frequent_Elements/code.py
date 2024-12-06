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

   freq = [[] for _ in range(maxoccur+1)]
   
   for i in hashMap:
      freq[hashMap[i]].append(i)

   while len(result) != k:
      for i in freq[maxoccur]:
         if len(result) != k:
            result.append(i)

      maxoccur -= 1

   return result

nums = [5,3,1,1,1,3,73,1]
print('Answer: ',topKFrequent(nums,2))