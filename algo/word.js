// 求出字符串 s 与字符串 t 的第⼆⻓公共单词(这⾥，
// 假设两个字符串均由英⽂字⺟和空格字符组成)；若找
// 到这样的公共单词，函数返回该单词，否则，函数返回
// NULL，如果有多个满⾜要求，则返回第⼀个单词。
// 例如：若 s= “This is C programming text”，
// t=“This is a text for C programming”，
// 则函数返回“This”。

// 分割字符串分别为map和arr
// 遍历arr，是否存在map中的key

// 空间复杂度O(n) 定义了两个长度为n的数组和一个map类型
// 时间复杂度O(n) 目前不考虑从map中做对比消耗的时间，只考虑for循环所用的时间
/**
 * @param {string} s
 * @param {string} t
 * @return {string}
 */
var getSecondWord = function(s, t) {
    // 判断字符串是否为空，或者是否包含两个单词
    var reg = /\s/;
    if (!s && !t && !reg.test(s) && !reg.test(t)) {
        return null;
    }
    // 基于空格分割s为一个数组     
    var arr = s.split(' ');

    // 基于空格分割t为一个对象
    var prev = 0
    var obj = new Object();
    for (var i = 0; i < t.length; i++) {
        if (t[i] == ' ') {
            var temp = t.substring(prev, i);
            obj[temp] = temp.length;
            prev = i + 1;
        }
    }
    // 最后剩余的一个参数
    var last = t.substring(prev, t.length);
    obj[last] = last.length

    // 判断map中是否存在对应的key, 若存在，插入数组中arr1
    var arr1 = [];
    for (var j = 0; j < arr.length; j++) {
        if (obj[arr[j]]) {
            arr1.push(arr[j]);
        }
    }

    if (arr1.length < 2) {
        return null;
    }

    // 排序arr1
    return arr1.map(i => ({ raw: i, len: i.length }))
        .sort((p, n) => n.len - p.len)
        .map(i => i.raw)[1]
}