/**
 * 工具
 */
layui.define(function (exports) {

    var utils = {
        /**
         * 是否重复数组
         * @param arr
         * @returns {boolean}
         */
        isRepeatArr: function (arr) {
            let hash = {};
            for (let i in arr) {
                if (hash[arr[i]]) {
                    return true;
                }
                hash[arr[i]] = true;
            }
            return false;
        },
        /**
         * 去除重复数组
         * @param array
         * @returns {*[]}
         */
        uniqueArr: function (array) {
            array.sort();
            var re = [array[0]];
            for (var i = 1; i < array.length; i++) {
                if (array[i] !== re[re.length - 1]) {
                    re.push(array[i]);
                }
            }
            return re;
        }
    }

    exports('utils', utils);
});

