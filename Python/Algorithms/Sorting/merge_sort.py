def merge_sort(nums):
   if len(nums) <= 1:
      return nums
   
   mid_index = len(nums) // 2
   left_nums = merge_sort(nums[:mid_index])
   right_nums = merge_sort(nums[mid_index:])
   sorted = []

   left_index = 0
   right_index = 0

   while left_index < len(left_nums) and right_index < len(right_nums):
      if left_nums[left_index] < right_nums[right_index]:
         sorted.append(left_nums[left_index])
         left_index += 1
      else:
         sorted.append(right_nums[right_index])
         right_index += 1

   sorted += left_nums[left_index:]
   sorted += right_nums[right_index:]

   return sorted      
   
numbers = [2, 6, 1, 8, 10, 4, 1]
print(merge_sort(numbers))