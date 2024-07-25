#PROBLEM:
"""
Given an expression string exp, write a program to examine whether the pairs and the orders of “{“, “}”, “(“, “)”, “[“, “]” are correct in the given expression.
"""
#SOLUTION:
def isReversed(bracket):
   if bracket == '[':
      return ']'
   if bracket == '{':
      return '}'
   if bracket == '(':
      return ')'
   

def isBracketsBalanced(bracketString):
   stack = []

   for bracket in bracketString:
      if isReversed(bracket) != None:
         stack.append(bracket)

      else:
         if not stack:
            return False
         
         current = stack.pop()
         reversed_bracket = isReversed(current)

         # print(f'Current: {current} -> Reversed: {reversed_bracket} -> Bracket: {bracket}')
         if reversed_bracket == None or reversed_bracket != bracket:
            return False
         
   if stack:
      return False
   return True
         

# print(isBracketsBalanced("{}"))
print(isBracketsBalanced("{[]}[]"))
