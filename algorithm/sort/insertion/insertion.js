// node insertion.js
/**
 * insertion sort
 * @param {Array} arr
 * @return {Array}
 */
function insertionSort(arr) {
    var len = arr.length;

    var preIndex, current;
    for (var i = 0; i < len; i++) {
        current = arr[i];
        preIndex = i - 1;

        while (preIndex >= 0 && arr[preIndex] > current) {
            arr[preIndex + 1] = arr[preIndex];
            preIndex--;
        }
        arr[preIndex + 1] = current;
    }

    return arr;
}

console.log(insertionSort([1, 3, 4, 2, 7]));