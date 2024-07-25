"""
Given an array of strings strs, group the anagrams together. You can return the answer in any order.

An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.

 

Example 1:

Input: strs = ["eat","tea","tan","ate","nat","bat"]
Output: [["bat"],["nat","tan"],["ate","eat","tea"]]
Example 2:

Input: strs = [""]
Output: [[""]]
Example 3:

Input: strs = ["a"]
Output: [["a"]]
"""

#SOLUTION

def group_alg(arr: list[str]):
   # CREATE RESULT LIST OF LISTS
   result = []
   # ITERATE THROUGH THE LIST
   for i, n in enumerate(arr):
      if i >= len(arr) - 1:
         return
      # CREATE CURRENT LIST
      curRes = [n]
      cur = n
      nxt = list(arr[i+1])
      # SELECT CURRENT ITEM AND PUSH TO CURRENT LIST
      for j in range(0, len(nxt)):
         if (nxt[j] in cur):
            print(cur, nxt, nxt[j])
            cur = arr[i + 1]
            print(cur)
      # cur.append(n)
      # print(cur, nxt)
      


strs = ["eat","tea","tan","ate","nat","bat"]
group_alg(strs)