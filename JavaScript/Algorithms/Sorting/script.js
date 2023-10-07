//PROBLEM:
//Sorting Algorithm without the use of in-built Javascript methods

//SOLUTION:

//Helper for function to return the index of the smallest value in an array
const minIndex = (arr) => {
   index = 0;

   for (let i = 0; i < arr.length; i++) {
      index = arr[i] < arr[index] ? i : index
   };

   return index;
};

const selectionSort = (arr) => {
   sortedArr = [];
   arrCopy = [...arr]

   for (let i = 0; i < arr.length; i++) {
      sortedArr.push(arrCopy[minIndex(arrCopy)]);
      arrCopy.splice(minIndex(arrCopy), 1)
   };
   
   return sortedArr
};

const numbers = [2, 6, 1, 8, 10, 4, 1];

console.log(selectionSort(numbers))
