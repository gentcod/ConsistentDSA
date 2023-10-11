const quickSort = (arr) => {
   if (arr.length <= 1) {
      return arr;
   }
   let pivot = arr[0]
   let arrLesser = []
   let arrGreater = []

   for (let i = 1; i < arr.length; i++) {
      if (arr[i] <= pivot) {
         arrLesser.push(arr[i]);
      }
      else {
         arrGreater.push(arr[i]);
      }
   }

   return [...quickSort(arrLesser), pivot, ...quickSort(arrGreater)]
}

const numbers = [2, 6, 1, 8, 10, 4, 1]
console.log(quickSort(numbers))
