#PROBLEM:
#Sorting Algorithm without the use of in-built Javascript methods

#SOLUTION:

#Helper for function to return the index of the smallest value in an array
def minIndex (arr):
   index = 0

   for i in range(len(arr)):
      if arr[i] < arr[index]: index = i

   return index

def selectionSort (arr):
   sortedArr = []

   for i in range(len(arr)):
      sortedArr.append(arr.pop(minIndex(arr)))
   
   return sortedArr

numbers = [2, 6, 1, 8, 10, 4, 1]

print(selectionSort(numbers))