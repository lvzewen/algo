// node selection.js
/**
 * selection sort
 * @param {Array} arr
 * @return {Array}
 */
function selectionSort(arr) {
    var len = arr.length;
    var minIndex, temp;

    for (var i = 0; i < len; i++) {
        minIndex = i;
        for (var j = i + 1; j < len; j++) {
            if (arr[minIndex] > arr[j]) {
                minIndex = j;
            }
        }
        temp = arr[i];
        arr[i] = arr[minIndex];
        arr[minIndex] = temp;
    }

    return arr;
}

console.log(selectionSort([1, 5, 3, 5, 3]))