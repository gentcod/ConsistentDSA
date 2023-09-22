""""Given two strings s and t, return true if s is a subsequence of t, or false otherwise.

A subsequence of a string is a new string that is formed from the original string by deleting some (can be none) of the characters 
without disturbing the relative positions of the remaining characters. (i.e., "ace" is a subsequence of "abcde" while "aec" is not)."""
def is_subsequence(s, t):
   ans = ''
   for i in range(len(t)):
      if s.find(t[i]) >= 0 and t[i] == s[len(ans)]:
         ans += t[i]
   print(ans)
   return ans == s

print(is_subsequence("abc", "ahbgdc"))
print(is_subsequence("acb", "ahbgdc"))
print(is_subsequence("ab", "baab"))
# print(is_subsequence("", ""))
