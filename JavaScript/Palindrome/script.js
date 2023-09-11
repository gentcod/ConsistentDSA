/**PROBLEM:
 * Given an integer x, return true if x is a palindrome, and false otherwise.
 */

//SOLUTION:
const isPalindrome = function(x) {
   let arr = x.toString().split('');
   let firstStr = arr.join('')
   let secondStr = arr.reverse().join('');
   return firstStr === secondStr
};

console.log(isPalindrome(212));