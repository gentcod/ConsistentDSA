const containsDuplicate = (nums) => {
   const unique = new Map();

   for (let i = 0; i < nums.length; i++) {
      if (unique.has(nums[i])) return true;
      unique.set(nums[i], true);
   } 
   return false;
}

console.log(containsDuplicate([1,2,3]));