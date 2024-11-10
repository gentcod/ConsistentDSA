# TIME COMPLEXITY - 0(n), SPACE COMPLEXITY - 0(1)
def encode(strs: list[str]) -> str:
   res = ""
   for s in strs:
      res += str(len(s)) + "#" + s
   return res

def decode(s: str) -> list[str]:
   res, i = [], 0

   while i < len(s):
      j = i
      while s[j] != "#":
         j += 1

      length = int(s[i:j])
      i = j + 1
      j = i + length
      res.append(s[i:j])
      i = j
   return res

# TIME COMPLEXITY - 0(n), SPACE COMPLEXITY - 0(N)
# def encode(strs: list[str]) -> str:
#    if not strs:
#       return ""
#    sizes, res = [], ""
#    for s in strs:
#       sizes.append(len(s))
#    for sz in sizes:
#       res += str(sz) + "_"

#    res += '#'
#    for s in strs:
#       res += s
#    return res

# def decode(s: str) -> list[str]:
#    res = []
#    arr = s.split("_")
#    chars = arr[-1][1:]

#    keys = arr[:-1]
#    for i, k in enumerate(keys):
#       res.append(chars[:int(k)])
#       chars = chars[int(k):]
#    return res

   
# pass_phrase = encode(["neet","code","love","you"])
# print(pass_phrase)

# print(decode(pass_phrase))