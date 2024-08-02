# def containsDuplicates(nums: list):
#    unique = set(nums)
#    return len(unique) != len(nums)

# Using hashMaps to short circuit return the moment a duplicate is located
def containsDuplicates(nums: list):
   unique = {}
   for i, n in enumerate(nums):
      if n in unique:
         return True
      unique[n] = i
   return False

print(containsDuplicates([1, 2, 3, 1]))