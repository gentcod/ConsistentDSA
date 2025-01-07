# Leetcode Problem: https://leetcode.com/problems/is-subsequence/

def is_subsequence(s, t):
   cur = s
   for i in range(len(t)):
      if len(s) == 0:
         return True
      if t[i] == cur[0]:
         cur = cur[1:]
   return len(cur) == 0

print(is_subsequence("abc", "ahbgdc"))
print(is_subsequence("acb", "ahbgdc"))
print(is_subsequence("ab", "baab"))
