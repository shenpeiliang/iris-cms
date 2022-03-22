/**
 * Created by shenpeiliang on 2019/12/24.
 */
$(function () {
    /**
     * [获取URL中的参数名及参数值的集合]
     * 示例URL:http://htmlJsTest/getrequest.html?uid=admin&rid=1&fid=2&name=小明
     * @param {[string]} urlStr [当该参数不为空的时候，则解析该url中的参数集合]
     * @return {[string]}       [参数集合]
     */
    function GetRequest(urlStr) {
        if (typeof urlStr == "undefined") {
            var url = decodeURI(location.search); //获取url中"?"符后的字符串
        } else {
            var url = "?" + urlStr.split("?")[1];
        }
        var theRequest = new Object();
        if (url.indexOf("?") != -1) {
            var str = url.substr(1);
            strs = str.split("&");
            for (var i = 0; i < strs.length; i++) {
                theRequest[strs[i].split("=")[0]] = decodeURI(strs[i].split("=")[1]);
            }
        }
        return theRequest;
    }

    var param = GetRequest()

    $.each(param, function (k, v) {
        //input
        if (v.length && $('input[name=' + k + ']').length) {
            $('input[name=' + k + ']').val(v)
        }

        //select
        if (v.length && $('select[name=' + k + ']').length) {
            $('select[name=' + k + ']').find("option[value=" + v + "]").attr('selected', true)
        }
    })

})

