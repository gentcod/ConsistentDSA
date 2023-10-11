def quicksort(arr):
   if len(arr) <= 1:
      return arr

   pivot = arr[0]
   arrLess = []
   arrGreater = []
   for num in arr[1:]:
      if num <= pivot:
         arrLess.append(num)
      else:
         arrGreater.append(num)
   
   return quicksort(arrLess) + [pivot] + quicksort(arrGreater)

# numbers = [2, 6, 1, 8, 10, 4, 1]
# print(quicksort(numbers))