/**
 * PROBLEM:
 * Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

   You may assume that each input would have exactly one solution, and you may not use the same element twice.

   You can return the answer in any order.
 */
//Crude
const twoSum = function(nums, target) {
   let index = [];
   for (let i of nums) {
      for (let j of nums) {
         if (nums.length <= 2 && i === j && i + j === target && index.length < 2) {
            return [0, 1];
         }
         if (i + j === target && index.length < 2) {
            index.push(nums.indexOf(j));
         }
      }
   }
   return index;
};

// console.log(twoSum([2,11,15,7],9));
// console.log(twoSum([3,2,3], 6));

//Optimized - Using HashMap
const twoSumOpt = function(nums, target) {
   let newMap = new Map();
   try {

      for (let i = 0; i < nums.length; i++) {      
         if (newMap.has(target - nums[i]))
         return [i, newMap.get(target - nums[i])]
         
         newMap.set(nums[i], i)
      }
   } finally {
      delete newMap;
   }
};

console.log(twoSumOpt([2,11,15,7],9));
console.log(twoSumOpt([3,2,3], 6));
