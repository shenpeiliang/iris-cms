/**
 * 活动管理列表
 */
!function () {
    layui.use(['form', 'ie_version'], function () {
        var form = layui.form;
        var ie_version = layui.ie_version;

        if (ie_version() < 10 && ie_version() != -1) {
            setTimeout(function () {
                $('#J_LoginBox').hide();

                layer.open({
                    title: '提示',
                    content: '您的浏览器过旧，请升级浏览器！',
                    end: function () {
                        location.href = 'https://www.google.cn/intl/zh-CN/chrome/';
                    }
                });
            }, 200);
        }

        //验证码
        $('#verify img').each(function () {
            var url = $(this).attr('src');
            $(this).click(function () {
                $(this).attr('src', url + '?randow=' + Math.random());
            });

        });

        //登录表单提交
        form.on('submit(login)', function (data) {
            layer.load();
            $.ajax({
                url: data.form.action,
                data: data.field,
                type: 'post',
                dataType: 'json',
                success: function (res) {
                    if (res.code == '200') {
                        layer.msg(res.msg, { icon: 6 }, function (index) {
                            layer.close(index)
                            if (res.data.url) {
                                window.location.href = res.data.url
                            }

                        });
                    } else {
                        layer.msg(res.msg, { icon: 5 }, function (index) {
                            layer.close(index)
                            $('#verify img').trigger('click')
                        });
                    }
                    layer.closeAll('loading');
                },
                error: function () {
                    layer.msg('系统错误', function () {
                        layer.closeAll('loading');
                    });
                }
            });
            return false;
        });
    });
}();