// node binary_tree.js

/**
 * binary tree sort
 * @param {Array} arr
 * @return {Array}
 */
function binaryTreeSort(arr) {
    var len = arr.length;
    if (len < 2) {
        return arr;
    }
    var tree = new BTS(arr[0]);

    for (var i = 1; i < len; i++) {
        tree.insert(arr[i]);
    }

    return tree.all();
}

function BTS(value = null, left = null, right = null) {
    this.left = left;
    this.value = value;
    this.right = right;
    this.insert = function(value) {
        if (!this.value) {
            this.value = value;
        }

        if (value < this.value) {
            if (!this.left) {
                this.left = new BTS(value);
            } else {
                this.left.insert(value);
            }
        } else {
            if (!this.right) {
                this.right = new BTS(value);
            } else {
                this.right.insert(value);
            }
        }
    };
    this.all = function() {
        var arr = [];
        return addValues(arr, this);
    };
}

// 将一个节点加入切片中
function addValues(values, t) {
    if (t) {
        values = addValues(values, t.left);
        values.push(t.value);
        values = addValues(values, t.right);
    }
    return values;
}

console.log(binaryTreeSort([1, 3, 4, 5, 10, 3]));