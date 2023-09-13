//PROBLEM:
//Write a function to find the longest common prefix string amongst an array of strings.
//If there is no common prefix return an empty string "".

//SOLUTION:
//Crude
const longestCommonPrefix = (arr) => {
   if (arr.every(str => str === arr[0])) return arr[0];
   let ans = '';
   let first = arr[0];
      
   for (let i = 0; i < first.length; i++) {
      if (arr.every(str => str[i] === first[i])) ans += first[i];
      else break;
   }

   return ans;
};

// console.log(longestCommonPrefix(["flower", "flow", "flight"]));
// console.log(longestCommonPrefix(["dog","racecar","car"]));
// console.log(longestCommonPrefix(["flower","flower","flower","flower"]))
// console.log(longestCommonPrefix(["babb","caa"]))
// console.log(longestCommonPrefix(['cir', 'car']))
// console.log(longestCommonPrefix(["abca","abc"]))