def intToRoman(n: int) -> str:
   result = ""
   keys = [1000, 500, 100, 50, 10, 5, 1]
   literals = {
		1000: "M",
		500: "D",
		100: "C",
		50: "L",
		10: "X",
		5: "V",
		1 : "I",
	}

   for i in range(len(keys)):
      if n % keys[i] != n:
         div = int((n - (n % keys[i]) ) / keys[i])
         if div == 4:
            if len(result) > 0 and result[-1] == literals.get(keys[i-1]):
               result = result[:-1]
               result += literals.get(keys[i]) + literals.get(keys[i-2])
            else:
               result += literals.get(keys[i]) + literals.get(keys[i-1])
         else:
            while div != 0:
               result += literals.get(keys[i])
               div -= 1
      n = n % keys[i]

   return result



print(intToRoman(4))