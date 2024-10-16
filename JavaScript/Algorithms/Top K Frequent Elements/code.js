const topKFrequentElements = (nums, k) => {
   hashMap = new Map()
   freq = []
   max = 0
   result = []

   for (let i = 0; i < nums.length; i++) {
      if (hashMap.has(nums[i])) hashMap.set(nums[i], (hashMap.get(nums[i])) + 1)
      else hashMap.set(nums[i], 1);

      if (max < hashMap.get(nums[i]))
         max = hashMap.get(nums[i])
   }

   for (let i = max; i >= 0; i--) {      
      freq.push([])
   }
   
   for (let i of hashMap) {
      const [val, index] = i;
      freq[index].push(val);
   }

   while (result.length != k) {
      for (let i of freq[max]) {         
         if (result.length != k) result.push(i)
      }
      max--
   }
   
   return result;
}

const nums = [5,3,1,1,1,3,73,1]
console.log(topKFrequentElements(nums, 2))