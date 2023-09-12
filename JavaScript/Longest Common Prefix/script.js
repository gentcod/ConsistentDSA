//PROBLEM:
//Write a function to find the longest common prefix string amongst an array of strings.
//If there is no common prefix return an empty string "".

//SOLUTION:
const longestCommonPrefix = (arr) => {
   if (arr.every(val => val === arr[0])) return arr[0];
   let ans = '';

   const convertStrToArr = (str) => str.split(''); 

   //Sort array based on length of string
   arr.sort((a, b) => a.length - b.length);

   let first = convertStrToArr(arr[0]);
      
   for (let i = 0; i < first.length; i++) {
      if (arr[i] !== undefined) convertStrToArr(arr[i]).includes(first[i]) && arr.every((cur) => convertStrToArr(cur)[i] === first[i]) ? ans += first[i] : '';
   }

   return ans;
};

// console.log(longestCommonPrefix(["flower", "flow", "flight"]));
// console.log(longestCommonPrefix(["dog","racecar","car"]));
console.log(longestCommonPrefix(["flower","flower","flower","flower"]))
// console.log(longestCommonPrefix(["babb","caa"]))