// node quick.js
/**
 * quick sort
 * @param {Array} arr
 * @param {number} left
 * @param {number} right
 * @return {Array}
 */
function quickSort(arr, left, right) {
    var len = arr.length,
        partitionIndex,
        left = typeof left != 'number' ? 0 : left,
        right = typeof right != 'number' ? len - 1 : right;

    if (left < right) {
        partitionIndex = partition(arr, left, right);
        quickSort(arr, left, partitionIndex - 1);
        quickSort(arr, partitionIndex + 1, right);
    }
    return arr;
}


/**
 * partition 分区操作
 * @param {Array} arr
 * @param {number} left
 * @param {number} right
 * @return {Array}
 */
function partition(arr, left, right) {
    var pivot = left, // 设定基准值（pivot）
        index = pivot + 1;
    for (var i = index; i <= right; i++) {
        if (arr[i] < arr[pivot]) {
            swap(arr, i, index);
            index++;
        }
    }
    swap(arr, pivot, index - 1);
    return index - 1;
}

/**
 * swap
 * @param {Array} arr
 * @param {number} i
 * @param {number} j
 * @return {Array}
 */
function swap(arr, i, j) {
    var temp = arr[i];
    arr[i] = arr[j];
    arr[j] = temp;
}


/**
 * partition
 * @param {Array} arr
 * @param {number} low
 * @param {number} high
 * @return {Array}
 */

function partition2(arr, low, high) {
    let pivot = arr[low];
    while (low < high) {
        while (low < high && arr[high] > pivot) {
            --high;
        }
        arr[low] = arr[high];
        while (low < high && arr[low] <= pivot) {
            ++low;
        }
        arr[high] = arr[low];
    }
    arr[low] = pivot;
    return low;
}

/**
 * quick sort
 * @param {Array} arr
 * @param {number} left
 * @param {number} right
 * @return {Array}
 */
function quickSort2(arr, low, high) {
    if (low < high) {
        let pivot = partition2(arr, low, high);
        quickSort2(arr, low, pivot - 1);
        quickSort2(arr, pivot + 1, high);
    }
    return arr;
}

console.log(quickSort([1, 22, 13, 34, 15], 0));
console.log(quickSort2([1, 22, 13, 34, 15], 0, 4));