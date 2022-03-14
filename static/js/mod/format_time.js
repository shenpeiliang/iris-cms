/**
 * 时间戳转日期
 */
layui.define(function (exports) {
    function formatNumber(n) {
        n = n.toString()
        return n[1] ? n : '0' + n;
    }

    function formatTime(number, format) {
        if (!format) {
            format = 'Y年M月D日 h:m:s';
        }
        number = number * 1000;
        var time = new Date(number)
        var newArr = []
        var formatArr = ['Y', 'M', 'D', 'h', 'm', 's']
        newArr.push(time.getFullYear())
        newArr.push(formatNumber(time.getMonth() + 1))
        newArr.push(formatNumber(time.getDate()))
        newArr.push(formatNumber(time.getHours()))
        newArr.push(formatNumber(time.getMinutes()))
        newArr.push(formatNumber(time.getSeconds()))

        for (var i in newArr) {
            format = format.replace(formatArr[i], newArr[i])
        }
        return format;
    }


    exports('format_time', formatTime);
});