//只能输入数字
function is_num_int (obj) {   // 值允许输入数字
    obj.value = obj.value.replace(/[^\d]/g, "");     //先把非数字的都替换掉，除了数字
    obj.value = obj.value.replace(/^0{0,}/, "");
}
//只能输入数字和一个小数点
function is_num (obj) {   // 值允许输入一个小数点和数字
    obj.value = obj.value.replace(/[^\d.]/g, "");     //先把非数字的都替换掉，除了数字和.
    obj.value = obj.value.replace(/^\./g, "");         //必须保证第一个为数字而不是.
    obj.value = obj.value.replace(/\.{2,}/g, ".");   //保证只有出现一个.而没有多个.
    obj.value = obj.value.replace(".", "$#$").replace(/\./g, "").replace("$#$", ".");    //保证.只出现一次，而不能出现两次以上
}
(function ($) {
    $.fn.select_radio = function (obj, end_obj) {
        obj.each(function () {
            $(this).click(function () {
                var v = $(this).attr('is_show');
                end_obj.val(v);
                $(this).parent().children('.r_selected').removeClass('r_selected');
                $(this).addClass('r_selected');
            });
        });

    }
})(jQuery);
$(function () {
    $().select_radio($('.radio_1'), $('#is_show'));//是否启用
    $().select_radio($('.radio_2'), $('#is_audit'));//是否认证
    //是否是多图
    $('.radio_imgs').each(function () {
        $(this).click(function () {
            var v = $(this).attr('flag');
            if (v == 1) {
                $('#tr_images').show();
            } else {
                $('#tr_images').hide();
            }
            $('#is_images').val(v);
            $(this).parent().children('.r_selected').removeClass('r_selected');
            $(this).addClass('r_selected');
        });
    });
    //验证码
    $('#verify img').each(function () {
        var url = $(this).attr('src');
        $(this).click(function () {
            $(this).attr('src', url + '?randow=' + Math.random());
        });

    });
    //折叠菜单
    $('.is_show').click(function () {
        var status = $(this).find('dd').slideToggle();
    });
    //表单提交
    $("#asubmit").click(function () {
        var obj = $(this).closest('form');
        $(obj).submit();
    });
    //所有表单提交前验证
    $('form').on('submit', function () {
        var tel_obj = $(this).find('.tel-format');
        var tel_obj_length = tel_obj.length;
        for (var i = 0; i < tel_obj_length; i++) {
            if ($(tel_obj[i]).val() == '' || !check_tel($(tel_obj[i]).val())) {
                layer.tips('手机号码格式有误！', $(tel_obj[i]), { tips: [2, '#FF5722'] });
                tel_obj[i].focus();
                return false;
            }
        }


        var obj = $('.validate');//必填
        var obj_length = obj.length;
        for (var i = 0; i < obj_length; i++) {
            if ($(obj[i]).val() == '') {
                $('html, body').animate({ scrollTop: 0 }, 'slow');
                if ($(obj[i]).attr('type') == 'hidden') {
                    layer.tips('必填项不能为空', $(obj[i]).parent());
                    return false;
                }
                layer.tips('必填项不能为空！', $(obj[i]));
                $(obj[i]).focus();
                return false;
            }
        }
        return true;
    });
    //生成微信菜单
    $('#create_menu').on('click', function () {
        $.post(create_menu, function (data) {
            if (data.length > 0) {
                layer.msg(data);
            } else {
                layer.msg('数据为空！');
            }
        });
    });
    //百度SEO主动推送
    $('#baidu_send').on('click', function () {
        var arr = $('.list_table').find('.bd_url');
        if (arr.length > 0) {
            $.post(baidu_send, function (data) {
                if (data.length > 0) {
                    layer.msg(data);
                } else {
                    layer.msg('数据为空！');
                }
            });
        } else {
            layer.msg('数据为空！');
        }

    });
    //全选
    $('#allChoose').each(function () {
        var default_value = $(this).attr("is_all");
        $(this).click(function () {
            if ($(this).attr("is_all") == default_value) {
                $('.checkitem').prop('checked', true);
                $(this).attr("is_all", "1");
            } else {
                $('.checkitem').prop('checked', false);
                $(this).attr("is_all", "0");
            }
        });

    });
    //发送邮件
    $('#btn_email').on('click', function () {
        var arr = "0";
        $('.checkitem:checked').each(function () {
            arr += ',' + $(this).val();
        });
        if (arr.length > 0) {
            $.post(email_url, { 'arr': arr }, function (data) {
                if (data.status == 1) {
                    window.location.href = data.res;
                } else {
                    layer.msg(data.res);
                }
            });
        } else {
            return;
        }

    });
    //图片删除
    $('#image_2').on('click', '.pic_del', function () {
        var parent_li = $(this).parent();//先获取父类元素
        var pic_id = $(this).attr('pic_id');

        if (!pic_id) {
            parent_li.remove();

            //删除原值
            var images_old = $('#images').val();
            images_old = images_old.split(',');

            var image_del = $(this).siblings('img').attr('src');
            for (var i in images_old) {
                if (images_old[i] == image_del)
                    images_old.splice(i, 1);
            }
            $('#images').val(images_old.join());

            return;
        }

        $.post(pic_del, { 'pic_id': pic_id }, function (data) {
            if (data == 'ok') {
                parent_li.remove();
            }
        });
    });
    //批量删除
    $('#batchDel').click(function () {
        var ids = [];

        $('.J_FCheckbox').each(function () {
            if ($(this).is(':checked')) {
                ids.push($(this).val());
            }
        });

        if (ids.length <= 0) {
            layer.msg('请选择要操作的记录', { icon: 5, time: 2000 });
            return;
        }

        $.post(del_url, { 'ids': ids }, function (e) {
            if (e.code == '200') {
                layer.msg(e.msg, function () {
                    window.location.reload()
                })
            } else {
                layer.msg(e.msg)
            }
        }, 'json');

    });
    //更新排序
    $('#updateOrd').click(function () {
        var json = {};
        $('.info').each(function () {
            key = $(this).find('.J_FCheckbox').val();
            val = $(this).find('.sort-input').val();
            json[key] = val;
        });

        JSON.stringify(json);
        if (json.length <= 0) {
            layer.msg('请选择要操作的记录', { icon: 5, time: 2000 });
            return;
        }

        $.post(order_url, { 'arr': json }, function (res) {
            if (res.code == '200') {
                location.reload();
            } else {
                layer.msg(res.msg, { icon: 5 });
            }
        }, 'json');

    });
    //验证管理员
    $("#admin_name").blur(function () {
        var user_name = $(this).val();
        $.post(validate_url, { 'user_name': user_name }, function (data) {
            if (data.length > 0) {
                $('#error').text(data);
            }
        });

    });
    //修改权限
    $(".auth_group").on('change', (function () {
        var group_id = $(this).val();
        $.post(validate_url, { 'group_id': group_id }, function (data) {
            if (data.length > 0) {
                layer.msg(data);
            }
        })
    })
    );
    //分类删除
    $('.del_family').on('click', function () {
        var url = $(this).attr('url');
        //询问框
        layer.confirm('删除父类将删除其下的所有子类？', {
            btn: ['确定', '取消'], //按钮
            shade: false //不显示遮罩
        }, function () {
            window.location.href = url;
        }, function () {
            var index = layer.confirm();
            layer.close(index);
        });

    });


    //模型变化
    $('#model').each(function () {
        $(this).change(function () {
            var model = $(this).val();
            if (model == '' || model == 'reply') {
                $('.tr_catid').hide();
                $('.tr_link').hide();
            } else if (model == 'def') {
                $('.tr_catid').hide();
                $('.tr_link').show();
            } else {
                $('.tr_catid').show();
                $('.tr_link').hide();
            }
            var uri = catid_url;
            var str = "<option value=''>请选择栏目</option>";
            $.post(uri, { 'model': model }, function (data) {
                if (data.length > 0) {
                    for (var i = 0; i < data.length; i++) {
                        str += "<option value='s_" + data[i].id + "'>" + data[i].title + "</option>";
                        if (data[i]._child) {
                            var child = data[i]._child;
                            if (child.length > 0) {
                                for (var j = 0; j < child.length; j++) {
                                    str += "<option value='" + child[j].id + "'>" + "&nbsp;&nbsp;&nbsp;" + child[j].title + "</option>";
                                }
                            }
                        }
                    }
                    $('#catid').html(str);
                }
            }, 'json');
        });

    });
    /*微信用户备注信息*/
    $('.member_remark').on('click', function () {
        var openid = $(this).attr('openid');
        var obj = $(this).parent('.handle').siblings('.txt_remark');
        layer.prompt({ title: '微信用户备注信息', formType: 0 }, function (text, index) {
            //layer.msg('您最后写下了：'+ text);
            $.post(remark_uri, { 'openid': openid, 'remark': text }, function (data) {
                if (data.length > 0) {
                    obj.html(text);
                    layer.close(index);
                    layer.msg(data, { icon: 1, time: 3000, title: '系统信息' });
                }
            });
        });

    });
    //数据库备份
    $(".btn_backup").on('click', function () {
        var uri = $(this).attr('uri');
        var load_index = layer.load(2);
        $.get(uri, function (data) {
            if (data == '200') {
                layer.close(load_index);
                window.location.reload();
            }
        });
    });
    //数据库还原
    $(".btn_rl").on('click', function () {
        var uri = $(this).attr('uri');
        var load_index = layer.load(2);
        $.get(uri, function (data) {
            if (data == '200') {
                layer.close(load_index);
                layer.msg('数据库还原成功！', { time: 5000, icon: 6 });
            }
        });
    });
    //其子类选择=模块
    $('.choose_1').each(function () {
        var default_value = $(this).attr("is_all");
        $(this).click(function () {
            if ($(this).attr("is_all") == default_value) {
                $(this).parents('dl').find('.checkitem').prop('checked', true);
                $(this).attr("is_all", "1");
            } else {
                $(this).parents('dl').find('.checkitem').prop('checked', false);
                $('#allChoose').prop('checked', false);
                $(this).attr("is_all", "0");
            }
        });
    });
    $('.choose_2').each(function () {
        var default_value = $(this).attr("is_all");
        $(this).click(function () {
            if ($(this).attr("is_all") == default_value) {
                $(this).parents('dd').find('.checkitem').prop('checked', true);
                $(this).attr("is_all", "1");
            } else {
                $(this).parents('dd').find('.checkitem').prop('checked', false);
                $('#allChoose').prop('checked', false);
                $(this).parents('dl').find('dt .checkitem').prop('checked', false);
                $(this).attr("is_all", "0");
            }
        });
    });
    //商品分类选择
    $('#set_sortid').each(function () {
        var obj_1 = $('#set_sortid select:eq(0)');
        var add_html = obj_1.html();
        var a, b = "";

        var sort_name = $('#sort_name');

        $(obj_1).change(function () {
            var parentid = $(this).val();
            a = $(this).find("option:selected").text();
            sort_name.val(a + "  " + b);
            var str = "<select name='sortid_1'";
            if (is_see != 1) {
                str += " class='validate'";
            }
            str += "><option value=''>请选择分类</option>";
            if (check_empty(parentid)) {
                $('#set_sortid select:eq(1)').remove();
                return false;
            }
            $.post(sort_uri, { 'parentid': parentid }, function (data) {
                if (data.length > 0) {
                    for (var i = 0; i < data.length; i++) {
                        str += "<option value='" + data[i].id + "'>" + data[i].title + "</option>";
                    }
                    str += "</select>";
                    if ($('#set_sortid select:eq(1)').length) {
                        $('#set_sortid select:eq(1)').remove();
                    }
                    $('#set_sortid').append(str);
                } else {
                    $('#set_sortid select:eq(1)').remove();
                    sort_name.val(a);
                }
            }, 'json');
        });
        $(document).on('change', '#set_sortid select:eq(1)', function () {
            b = $(this).find("option:selected").text();
            sort_name.val(a + "  " + b);
            var c = "";
            var sort_id = $(this).find("option:selected").val();
            var str = "<select name='sortid_1' class='validate'>";
            $.post(sort_active_uri, { 'parentid': sort_id }, function (data) {
                if (data.length > 0) {
                    for (var i = 0; i < data.length; i++) {
                        str += "<option value='" + data[i][1] + "'>" + data[i][0] + "</option>";
                        c += data[i][1] + ":" + data[i][0] + ",";
                    }
                    str += "</select>";
                    $('.select_spec').html(str);
                    c = c.substr(0, c.length - 1);
                    $('#specification_name').val(c);
                }
            }, 'json');
        });

    });

    //配置
    $('.con_add').on('click', function () {
        var str = '<div class="specs"><div class="specs_1"><input type="text"  name="powers[title][]"  class="txt_basic validate" size="10" "/><i class="require-red">-</i><input type="text"  name="powers[name][]" class="txt_basic validate" size="10" /></div><div class="specs_2"><a href="javasript:void(0);" title="删除" class="con_del"></a></div></div>';
        $('#spec_set').append(str);
    });
    $('#spec_set').on('click', '.con_del', function () {
        if ($('#spec_set').find('.specs').length > 1) {
            var obj = $(this).parents('.specs');
            //数据删除
            obj.remove();
        }
    });
    //规格开启、关闭
    $('.btn_spec').each(function () {
        $(this).click(function () {
            var is_active = $('#is_active').val();
            if (is_active == 0) {
                $('#is_active').val(1);
                $('#good_spec').show();
                $(this).text('关闭规格');
            } else {
                $('#is_active').val(0);
                $('#good_spec').hide();
                $(this).text('开启规格');
            }
        });

    });
    //添加规格
    $('.spec_add').on('click', function () {
        var split_line = ' <i class="require-red">-</i> ';

        var str = '<div class="specs"><div class="specs_1">';
        str += '<input type="text"  name="spec_name[]"  class="txt_basic validate" size="8"/>';
        str += split_line;
        str += '<input type="text"  name="spec_price[]"  class="txt_basic validate" size="5" onkeyup="is_num(this)"/>';
        str += split_line;
        str += '<input type="text"  name="spec_stock[]"  class="txt_basic validate" size="5" onkeyup="is_num_int(this)"/>';
        str += split_line;
        str += '<input type="text"  name="spec_no[]"  class="txt_basic" size="8" />';
        str += split_line;
        str += '<input type="text"  name="spec_paixu[]"  class="txt_basic validate" size="3" value="100" onkeyup="is_num_int(this)"/>';
        str += split_line;
        str += '<input type="file"  name="spec_img[]"  class="txt_basic" size="3" /></div>';
        str += '<div class="specs_2"><a href="javascript:void(0);"  class="con_normal con_del">删除</a></div></div>';
        $('#spec_set').append(str);
    });
    //是否为空
    function check_empty (obj) {
        if (typeof (obj) == "undefined" || obj <= 0) {
            return true;
        } else {
            return false;
        }
    }

    //工匠信息
    $("#worker_uname_txt").on('keyup', (function () {
        var uname = $(this).val();
        if (uname.length == 0) {
            return;
        }
        var str = '';
        $.post(worker_url, { 'uname': uname }, function (data) {
            if (data.length > 0) {
                for (var i = 0; i < data.length; i++) {
                    str += '<li><a href="javascript:void(0);" uid=' + data[i].id + ' class="worker_info">' + data[i].title + '</a></li>';
                }
                $('#worker_lists ul').html(str);
                $('#worker_lists').show();
            }
        }, 'json');
    })
    );
    $("#worker_lists_show").on('click', '.worker_info', function () {
        $('#worker_uname_txt').val($(this).text());
        $('#uid').val($(this).attr('uid'));
        $('#worker_lists').hide();
    });
    //选择城市
    $('#city_to_btn').click(function () {
        var obj = $(this);
        var is_show = obj.attr('is_show');
        if (is_show == 1) {
            $('#city_to_con').hide();
            obj.attr('is_show', 0).text('展开城市');
        } else {
            $('#city_to_con').show();
            obj.attr('is_show', 1).text('折叠城市');
        }
        //$('#city_to_con').toggle();
    });

    //表单提交
    $('#form-submit').click(function () {
        //必填
        var obj = $('.validate');
        var obj_length = obj.length;
        for (var i = 0; i < obj_length; i++) {
            if ($(obj[i]).val() == '') {
                $('html, body').animate({ scrollTop: 0 }, 'slow');
                if ($(obj[i]).attr('type') == 'hidden') {
                    layer.tips('必填项不能为空', $(obj[i]).parent());
                    return false;
                }
                layer.tips('必填项不能为空！', $(obj[i]));
                $(obj[i]).focus();
                return false;
            }
        }

        var data = $("form").serialize()
        var url = $("form").attr('action')

        $.post(url, data, function (res) {
            if (res.code == '200')
                return layer.msg(res.msg, function () {
                    if (res.data.url)
                        return window.location.href = res.data.url
                    else
                        return window.location.reload()
                })
            else
                return layer.msg(res.msg)
        }, 'json')
    })

    //添加、修改弹窗方式
    $('.layer-form').click(function () {
        var html_url = $(this).attr('data-uri');
        var layer_title = $(this).text();

        layer.open({
            type: 2,
            area: ['1000px', '600px'],
            title: layer_title,
            content: html_url,
            btn: ['确定', '取消'],
            btn1: function (index) {
                var ifram_body = layer.getChildFrame('body', index);
                var form = $(ifram_body).find('form');

                //必填
                var obj = $(ifram_body).find('.validate');
                var obj_length = obj.length;
                for (var i = 0; i < obj_length; i++) {
                    if ($(obj[i]).val() == '') {
                        $('html, body').animate({ scrollTop: 0 }, 'slow');
                        if ($(obj[i]).attr('type') == 'hidden') {
                            layer.tips('必填项不能为空', $(obj[i]).parent());
                            return false;
                        }
                        layer.tips('必填项不能为空！', $(obj[i]));
                        $(obj[i]).focus();
                        return false;
                    }
                }

                var data = $(ifram_body).find("form").serialize()
                var url = $(ifram_body).find("form").attr('action')

                $.post(url, data, function (res) {
                    if (res.code == '200')
                        return layer.msg(res.msg, function () {
                            layer.close(index)
                            return window.location.reload()
                        })
                    else
                        return layer.msg(res.msg)
                }, 'json')
            },
            btn2: function (index) {
                layer.close(index)
            }
        })
    })

});

$(function () {

    layui.use(['laydate', 'upload', 'form', 'element', 'layedit'], function () {
        var laydate = layui.laydate;
        var upload = layui.upload;
        var form = layui.form;
        var element = layui.element;
        var layedit = layui.layedit;


        var isNum = function (v) {
            return /^[0-9]*$/.test(v);
        }

        form.render();

        form.verify({
            money: function (value) {
                if (value && !$.trim(value).match(/^[0-9]\d*(\.\d+)?$/)) {
                    return '价格为数字。填写中文、字母或特殊符号无效';
                }
            },
            toncat_qq: function (value) {
                if (!isNum(value) || value.length > 20) {
                    return '仅为数字且上限为20位字符';
                }
            },
            toncat_weixin: function (value) {
                if (value.length > 20) {
                    return '上限为20位字符';
                }
            },
            toncat_else: function (value) {
                if (value.length > 25) {
                    return '上限为25位字符';
                }
            },
            rich_edit: function (value) {
                //富文本
                if ($("textarea[data-layedit_index]").length > 0) {
                    $("textarea[data-layedit_index]").each(function (i) {
                        layedit.sync($("textarea[data-layedit_index]").eq(i).attr('data-layedit_index'));
                    });
                }
            }
        });

        laydate.render({
            elem: '#start',
            format: 'yyyy-MM-dd HH:mm:ss',
            type: 'datetime',
            done: function (value) {
                if (new Date($('#end').val()) < new Date(value)) {
                    layer.msg('开始时间不能晚于结束时间', { icon: 5, time: 2000 });
                    $('#start').val('');
                }
            }
        });

        laydate.render({
            elem: '#date',
        });

        laydate.render({
            elem: '#end',
            type: 'datetime',
            done: function (value) {
                if (new Date($('#start').val()) > new Date(value)) {
                    layer.msg('结束时间不能早于开始时间', { icon: 5, time: 2000 });
                    $('#end').val('');
                }
            }
        });


        //页面加载完成后重新渲染
        $(document).ready(function () {
            layui.form.render();
        });


        //点评等级
        form.on('radio(grade-select)', function (data) {
            grade = data.value;
            tag_type = (grade == 1 || grade == 2) ? 'bad' : 'good';
            var tag_select = 'remark-tag-option-' + tag_type;

            $('.remark-tag').show();

            var tags = $('.remark-tag div');

            for (var i = 0; i < tags.length; i++) {
                if ($(tags[i]).hasClass(tag_select)) {
                    $(tags[i]).show();
                } else {
                    $(tags[i]).hide();
                }
                form.render('checkbox');
            }
        });


        //权限选择
        form.on("checkbox(rule_checkbox)", function (data) {
            var $this = $(this);

            if (data.elem.checked) {
                $this.parents('dt').siblings('dd').find('input').prop('checked', true);
                form.render('checkbox');
            } else {
                $this.parents('dt').siblings('dd').find('input').prop('checked', false);
                form.render('checkbox');
            }
        });

        form.on("checkbox(rule_checkbox_div)", function (data) {
            var $this = $(this);

            if (data.elem.checked) {
                $this.parents('span').siblings('div').find('input').prop('checked', true);
                form.render('checkbox');
            } else {
                $this.parents('span').siblings('div').find('input').prop('checked', false);
                form.render('checkbox');
            }
        });

        //帖子类型选择
        form.on('radio(post_type_even)', function (data) {
            $('.post-section-main').removeClass('section-show').eq($(this).val() - 1).addClass('section-show');
        });

        //标签显示
        form.on('radio(is_multicast_event)', function (data) {
            if (data.value == 1) {
                $('.section-tags').removeClass('hide')
            } else {
                $('.section-tags').addClass('hide')
            }
        });

        //全选
        form.on("checkbox(check_all_event)", function (data) {
            $(this).parents('.checkbox-section').find('.check_item').prop('checked', data.elem.checked);
            form.render('checkbox');
        });


        //地区选择--省变化
        form.on('select(province)', function (data) {
            var parentid = data.value;
            var uri = linkage_url;
            var str = "<option value=''>请选择城市</option>";
            $.post(uri, { 'parentid': parentid }, function (ret) {
                if (ret.length > 0) {
                    for (var i = 0; i < ret.length; i++) {
                        str += "<option value='" + ret[i].id + "'>" + ret[i].name + "</option>";
                    }
                    $('#city').html(str);
                    form.render('select');
                }
            }, 'json');

        });
        //地区选择--城市变化
        form.on('select(city)', function (data) {
            var parentid = data.value;
            var uri = linkage_url;
            var str = "<option value=''>请选择区/县</option>";
            $.post(uri, { 'parentid': parentid }, function (ret) {
                if (ret.length > 0) {
                    for (var i = 0; i < ret.length; i++) {
                        str += "<option value='" + ret[i].id + "'>" + ret[i].name + "</option>";
                    }
                    $('#county').html(str);
                    form.render('select');
                }
            }, 'json');

            //地区
            if ($('#village').length) {
                village(parentid, $('#county').find("option:selected").val());
            }
        });

        //区县
        form.on('select(county)', function (data) {
            var parentid = data.value;

            //地区
            if ($('#village').length) {
                village($('#city').find("option:selected").val(), parentid);
            }
        });

        //小区
        var village = function (city, county) {
            var city = city || 0;
            var county = county || 0;

            var uri = "/tool/main/get_village.html";
            var str = "<option value=''>请选择所属小区</option>";
            $.post(uri, {
                'city': city,
                'county': county
            }, function (ret) {
                if (ret.length > 0) {
                    for (var i = 0; i < ret.length; i++) {
                        str += "<option value='" + ret[i].id + "'>" + ret[i].title + "</option>";
                    }
                }
                $('#village').html(str);
                form.render('select');
            }, 'json');
        };

        var layedit_init = function (id) {
            layui.use('layedit', function () {
                var layedit = layui.layedit;
                var index = layedit.build(id, {
                    height: '200px',
                    tool: [
                        'face' //表情
                    ],
                    uploadImage: {
                        url: '/game/uploadify/layedit?img_type=forum'
                    }
                });


                $('#' + id).attr('data-layedit_index', index);
            });
        };

        $('.layedit').each(function () {
            layedit_init($(this).attr('id'));
        });
        //表单提交
        form.on('submit(saveBtn)', function (data) {
            var uri = $('.layui-form').attr('action');

            admin.ajax({
                url: uri,
                type: 'post',
                data: data.field
            }).done(function (res) {
                if (res.code != '200') {
                    layer.msg(res.msg || '操作失败', { icon: 5, time: 2000 });
                    return;
                }

                layer.msg(res.msg || "操作成功", { icon: 6, time: 1000 }, function () {
                    if (window != top) {
                        top.location.reload();
                        return;
                    }
                    window.location.reload()
                });
            });
            return false;
        });

        //上传图片
        $('.J_UploadImgBtn').each(function () {
            var $this = $(this);
            upload.render({
                elem: '#' + $this.attr('id'),
                url: '/admin/uploadify/upload?img_type=' + $this.attr('data-itype'),
                size: 1024,
                field: 'Filedata',
                acceptMime: 'image/*',
                before: function () {
                    layer.load();
                },
                done: function (res) {
                    layer.closeAll('loading');

                    if (res.code == '200') {
                        $this.siblings('img').removeClass('hide').attr({ 'src': res.data.path }).data('img', res.data.path);
                        $this.siblings('input[type="hidden"]').val(res.data.path);
                    } else {
                        layer.msg(res.msg || '操作失败', { icon: 5, time: 2000 });
                    }
                },
                error: function () {
                    layer.closeAll('loading');
                    layer.msg('网络错误', { icon: 5, time: 2000 });
                }
            });
        });

        //多图上传
        $('.J_UploadImgMoreBtn').each(function () {
            var $this = $(this);
            upload.render({
                elem: '#' + $this.attr('id'),
                multiple: true,
                url: '/admin/uploadify/upload?img_type=' + $this.attr('data-itype'),
                size: 1024,
                field: 'Filedata',
                acceptMime: 'image/*',
                before: function () {
                    layer.load();
                },
                done: function (res) {
                    layer.closeAll('loading');
                    var _img = '';
                    if (res.code == '200') {
                        _img += '' +
                            '<span class="J_ImgItem">' +
                            '   <input type="hidden" name="images[]" value="' + res.data + '"/>' +
                            '   <a href="javascript:;" class="J_ViewImg" data-img="' + res.data + '"><img  class="upload-img" src="' + res.data + '" /></a>' +
                            '   <a href="javascript:;" data-img="' + res.data + '" class="layui-icon layui-icon-close del J_Del"></a>' +
                            '</span>';
                        $('#J_Pzbox').append(_img);

                    } else {
                        layer.msg(res.msg || '操作失败', { icon: 5, time: 2000 });
                    }
                },
                error: function () {
                    layer.closeAll('loading');
                    layer.msg('网络错误', { icon: 5, time: 2000 });
                }
            });
        });

        //删除操作
        $('.event-delete').click(function () {
            $.get($(this).attr('data-uri'), function (e) {
                if (e.code != '200') {
                    layer.msg(e.msg);
                    return;
                }
                layer.msg(e.msg, { time: 2000 }, function () {
                    location.reload()
                })
            }, 'json')
        });

        //查看操作记录
        $('.show-log').on('click', function () {
            var $this = $(this);
            var uri = $this.data('uri');

            xadmin.open('执行记录', uri);
            form.render();
        });

        //用户表单
        $('.user-form').on('click', function () {
            var $this = $(this);
            var uri = $this.data('uri');

            xadmin.open('用户信息', uri);
            form.render();
        });

        //排序
        function getUrlParam (url, name) {
            var reg = new RegExp('(^|&)' + name + '=([^&]*)(&|$)');
            var len = url.indexOf('?') + 1;
            var query = url.substring(len);
            var result = query.match(reg);
            return result ? decodeURIComponent(result[2]) : null;
        }

        var resultUrl = window.location.href;
        var $asc = getUrlParam(resultUrl, "asc");
        var ascFuc = function () {
            if ($asc == "false" || $asc == null) {
                $asc = "true"
            } else {
                $asc = "false"
            }
        };
        $(".orderby-span").each(function (i) {
            var $this = $(this);

            $this.click(function () {
                if (i == 0 || i == 1) {
                    ascFuc()
                }

                $keyword_type = getUrlParam(resultUrl, "keyword_type");
                $keyword = getUrlParam(resultUrl, "keyword");

                var uri = $this.data('uri');
                uri += '&keyword_type=' + $keyword_type;
                uri += '&keyword=' + $keyword;
                uri += '&sort=' + i;
                uri += '&asc=' + $asc;
                window.location.href = uri;
            })
        });

        //删除所选
        $('#event-delall').on('click', function () {
            var ids = [];
            var uri = $(this).data('uri');

            $('.J_FCheckbox').each(function () {
                var $this = $(this);
                if ($this.is(':checked')) {
                    ids.push($(this).val());
                }
            });

            if (ids.length <= 0) {
                layer.msg('请选择要操作的记录', { icon: 5, time: 2000 });
                return;
            }

            layer.confirm('确定要删除吗？', function () {
                admin.ajax({
                    url: uri,
                    data: { ids: ids },
                    type: 'post'
                }).done(function (res) {
                    if (res.code != 'ok') {
                        layer.msg(res.msg || '操作失败', { icon: 5, time: 2000 });
                        return;
                    }
                    layer.msg(res.msg || "操作成功", { icon: 6, time: 2000 }, function () {
                        location.reload();
                    });
                });
            });
        });

        /**
         * 获取基础信息
         * @param url 请求url
         * @param callback
         */
        var getBaseOptions = function (url, callback) {
            admin.ajax({
                url: url,
                type: 'post'
            }).done(function (res) {
                if (res.success) {
                    callback(res.msg);
                } else {
                    layer.msg(res.msg || '获取数据失败', { icon: 5, time: 2000 });
                }
            });
        };


        $('.post-relation-form').on('click', function () {
            var uri = $(this).data('uri');

            layer.open({
                title: '表单信息',
                area: ['420px', '240px'],
                type: 1,
                content: '' +
                    '<form class="layui-form dialog-form" action="" id="J_Form">\n' +
                    '  <div class="dialog-form-item">' +
                    '                        <label class="dialog-form-label">授权类型</label>' +
                    '                        <div class="dialog-form-input">' +
                    '                            <input type="radio" lay-filter="rtype_even" name="r_type" value="0" checked title="指定用户" />' +
                    '                            <input type="radio" lay-filter="rtype_even" name="r_type" value="1" title="全部用户" />' +
                    '                        </div>' +
                    '                    </div>' +
                    '    <div class="dialog-form-ite" id="form-ids">\n' +
                    '        <label class="dialog-form-label">用户ID</label>\n' +
                    '        <div class="dialog-form-input">\n' +
                    '            <input type="text" id="J_Title" name="ids" placeholder="请输入用户ID" maxlength="10" autocomplete="off" value="">\n' +
                    '            <p style="color:#999">多个用户ID使用英文逗号隔开</p>' +
                    '        </div>\n' +
                    '    </div>\n' +
                    '    <div class="dialog-form-item">\n' +
                    '        <label class="dialog-form-label">&nbsp;</label>\n' +
                    '        <div class="dialog-form-input">\n' +
                    '            <a href="javascript:;" id="J_Btn" class="dialog-form-btn">确定</a>\n' +
                    '            <a href="javascript:layer.closeAll();" class="dialog-form-hollowbtn">取消</a>\n' +
                    '        </div>\n' +
                    '    </div>\n' +
                    '</form>',
                success: function () {
                    form.render('radio');
                    form.on('radio(rtype_even)', function (data) {
                        if (data.value == 1) {
                            $('#form-ids').hide();
                        } else {
                            $('#form-ids').show();
                        }
                    });
                    $('#J_Btn').on('click', function () {

                        var title = $.trim($('#J_Title').val());
                        var stype = $("input[name='r_type']:checked").val();
                        console.log(stype)
                        if (stype == 0 && !title) {
                            layer.msg('请输入用户ID', { icon: 5, time: 2000 });
                            return;
                        }

                        admin.ajax({
                            url: uri,
                            type: 'post',
                            data: {
                                r_type: stype,
                                ids: $('#J_Title').val()
                            },
                        }).done(function (res) {
                            if (res.code != '200') {
                                layer.msg(res.msg || '操作失败', { icon: 5, time: 2000 });
                                return;
                            }
                            layer.alert("保存成功", { icon: 6 }, function () {
                                layer.closeAll();
                            });
                        });
                    });
                }
            });
        });

        //图文关联店铺
        $('.news-relation-form').on('click', function () {
            var uri = $(this).data('uri');

            layer.open({
                title: '表单信息',
                area: ['420px', '240px'],
                type: 1,
                content: '' +
                    '<form class="layui-form dialog-form" action="" id="J_Form">\n' +
                    '  <div class="dialog-form-item">' +
                    '                        <label class="dialog-form-label">关联类型</label>' +
                    '                        <div class="dialog-form-input">' +
                    '                            <input type="radio" lay-filter="rtype_even" name="r_type" value="0" checked title="指定门店" />' +
                    '                            <input type="radio" lay-filter="rtype_even" name="r_type" value="1" title="全部门店" />' +
                    '                        </div>' +
                    '                    </div>' +
                    '    <div class="dialog-form-ite" id="form-ids">\n' +
                    '        <label class="dialog-form-label">店铺ID</label>\n' +
                    '        <div class="dialog-form-input">\n' +
                    '            <input type="text" id="J_Title" name="ids" placeholder="请输入店铺ID" maxlength="10" autocomplete="off" value="">\n' +
                    '            <p style="color:#999">多个店铺ID使用英文逗号隔开</p>' +
                    '        </div>\n' +
                    '    </div>\n' +
                    '    <div class="dialog-form-item">\n' +
                    '        <label class="dialog-form-label">&nbsp;</label>\n' +
                    '        <div class="dialog-form-input">\n' +
                    '            <a href="javascript:;" id="J_Btn" class="dialog-form-btn">确定</a>\n' +
                    '            <a href="javascript:layer.closeAll();" class="dialog-form-hollowbtn">取消</a>\n' +
                    '        </div>\n' +
                    '    </div>\n' +
                    '</form>',
                success: function () {
                    form.render('radio');
                    form.on('radio(rtype_even)', function (data) {
                        if (data.value == 1) {
                            $('#form-ids').hide();
                        } else {
                            $('#form-ids').show();
                        }
                    });
                    $('#J_Btn').on('click', function () {

                        var title = $.trim($('#J_Title').val());
                        var stype = $("input[name='r_type']:checked").val();
                        console.log(stype)
                        if (stype == 0 && !title) {
                            layer.msg('请输入店铺ID', { icon: 5, time: 2000 });
                            return;
                        }

                        admin.ajax({
                            url: uri,
                            type: 'post',
                            data: {
                                r_type: stype,
                                ids: $('#J_Title').val()
                            },
                        }).done(function (res) {
                            if (res.code != '200') {
                                layer.msg(res.msg || '操作失败', { icon: 5, time: 2000 });
                                return;
                            }
                            layer.alert("保存成功", { icon: 6 }, function () {
                                layer.closeAll();
                            });
                        });
                    });
                }
            });
        });

        //统一修改
        $('.event-form').on('click', function () {
            var uri = $(this).data('uri');
            var page_title = $(this).data('title') || '表单信息';
            xadmin.open(page_title, uri);
        });

        //活动类型选择
        form.on('radio(activity_type)', function (data) {
            if (data.value == 1) {
                $('.activity-packege-content').removeClass('hide')
            } else {
                $('.activity-packege-content').addClass('hide')
            }
        });

        //用户屏蔽
        $('.block-user').on('click', function () {
            var uid = $(this).data('uid');
            var uri = $(this).data('uri');
            var is_blocked = $(this).data('is_blocked');

            layer.confirm('确定要对用户操作' + $(this).text() + '吗？', function () {
                admin.ajax({
                    url: uri,
                    data: {
                        uid: uid,
                        is_blocked: is_blocked
                    },
                    type: 'post'
                }).done(function (res) {
                    if (res.code != '200') {
                        layer.msg(res.msg || '操作失败', { icon: 5, time: 2000 });
                        return;
                    }
                    layer.msg(res.msg || "操作成功", { icon: 6, time: 2000 }, function () {
                        location.reload();
                    });
                });
            });
        });

        //显示状态
        form.on('switch(state)', function (data) {
            var id = $(this).data('id');

            var url = $(this).data('url');
            var request_url = url || state_url;

            var val = data.elem.checked;
            admin.ajax({
                url: request_url,
                type: 'post',
                data: {
                    id: id,
                    is_show: this.checked ? 1 : 0
                }
            }).done(function (res) {
                if (res.code != '200') {
                    layer.msg(res.msg || '操作失败', { icon: 5, time: 2000 });
                    data.elem.checked = !val;
                    form.render();
                    return;
                }
                layer.msg("保存成功", { icon: 6, time: 1000 });
            });
        });

        //用户认证
        form.on('switch(auth_user)', function (data) {
            var id = $(this).data('id');
            admin.ajax({
                url: user_auth_uri,
                type: 'post',
                data: {
                    id: id,
                    is_auth: this.checked ? 1 : 0
                }
            }).done(function (res) {
                if (res.code != '200') {
                    layer.msg(res.msg || '操作失败', { icon: 5, time: 2000 });
                    data.elem.checked = !val;
                    form.render();
                    return;
                }
                layer.msg("保存成功", { icon: 6, time: 1000 });
            });
        });

        //是否推荐
        form.on('switch(recommend)', function (data) {
            var id = $(this).data('id');
            admin.ajax({
                url: recommend_uri,
                type: 'post',
                data: {
                    id: id,
                    is_recommend: this.checked ? 1 : 0
                }
            }).done(function (res) {
                if (res.code != '200') {
                    layer.msg(res.msg || '操作失败', { icon: 5, time: 2000 });
                    return;
                }
                layer.msg("保存成功", { icon: 6, time: 1000 });
            });
        });

        //查看大图
        $('body').on('click', '.J_ViewImg', function () {
            var $this = $(this);
            var img = '';
            if ($this.data('img').indexOf('.com') > 0) {
                img = $this.data('img');
            } else {
                img = $this.data('img');
            }
            layer.open({
                title: '查看大图',
                type: 1,
                content: '<div><img src="' + img + '" style="width:100%;display: block;"/></div>'
            });
        });

        //删除图片
        $('body').on('click', '.J_Del', function () {
            $(this).closest('.J_ImgItem').remove();
            $('#J_Upload').removeAttr('disabled');
        });

        //商品基准价格
        $('#price-config .goods-btn-add').click(function () {
            var _html = '<div class="layui-col-xs12 price-section-item">' +
                '<span class="layui-form-mid layui-word-aux">大于等于</span>' +
                '<div class="layui-input-inline layui-input-short">' +
                '<input name="price_json[min_num][]" type="number" maxlength="30" placeholder="请输入数量" lay-verify="required"' +
                '   lay-reqText="请输入数量" value="" class="layui-input">' +
                '</div>' +
                '<span class="layui-form-mid layui-word-aux">个，小于等于</span>' +
                '<div class="layui-input-inline layui-input-short">' +
                '<input name="price_json[max_num][]" type="number" maxlength="30" placeholder="请输入数量" lay-verify="required"' +
                '   lay-reqText="请输入数量" value="<{$data.min_price}>" class="layui-input">' +
                '</div>' +
                '<span class="layui-form-mid layui-word-aux">个</span>' +
                '<div class="layui-input-inline layui-input-short">' +
                '<input name="price_json[price][]" type="number" maxlength="30" placeholder="请输入价格" lay-verify="required"' +
                '   lay-reqText="请输入价格" value="<{$data.min_price}>" class="layui-input">' +
                '</div>' +
                '<span class="layui-form-mid layui-word-aux">元/个</span>' +
                '<button type="button" class="layui-btn layui-btn-sm goods-btn-delete"><i class="layui-icon">&#xe640;</i></button>' +
                '</div>';

            $(this).parents('#price-config').find('.price-section').append(_html);
        });

        $('#price-config').on('click', '.goods-btn-delete', function () {
            $(this).parents('.price-section-item').remove();
        });

        //活动推广
        $('.activity-union-form').on('click', function () {
            var uri = $(this).data('uri');
            var sid = $(this).data('sid') || '';
            var src_url = $(this).data('src_url') || '';

            layer.open({
                title: '表单信息',
                area: ['520px', '240px'],
                type: 1,
                content: '' +
                    '<form class="layui-form dialog-form" action="" id="J_Form">\n' +
                    '    <div class="dialog-form-ite" id="form-store_id" style="width: 100%;height:40px;">\n' +
                    '        <label class="dialog-form-label">店铺编号</label>\n' +
                    '        <div class="dialog-form-input">\n' +
                    '            <input type="text" id="J_store_id" name="sid" placeholder="请输入店铺编号" style="width: 96px;" autocomplete="off" value="' + sid + '">\n' +
                    '        </div>\n' +
                    '    </div>\n' +
                    '    <div class="dialog-form-ite" id="form-ids">\n' +
                    '        <label class="dialog-form-label">活动链接</label>\n' +
                    '        <div class="dialog-form-input">\n' +
                    '            <input type="text" id="J_Title" name="url" placeholder="请输入活动链接" style="width: 360px;" autocomplete="off" value="' + src_url + '">\n' +
                    '        </div>\n' +
                    '    </div>\n' +
                    '    <div class="dialog-form-item">\n' +
                    '        <label class="dialog-form-label">&nbsp;</label>\n' +
                    '        <div class="dialog-form-input">\n' +
                    '            <a href="javascript:;" id="J_Btn" class="dialog-form-btn">确定</a>\n' +
                    '            <a href="javascript:layer.closeAll();" class="dialog-form-hollowbtn">取消</a>\n' +
                    '        </div>\n' +
                    '    </div>\n' +
                    '</form>',
                success: function () {
                    $('#J_Btn').on('click', function () {

                        var title = $.trim($('#J_Title').val());
                        if (!title) {
                            layer.msg('请输入活动链接', { icon: 5, time: 2000 });
                            return;
                        }

                        var store_id = $.trim($('#J_store_id').val());
                        if (!store_id) {
                            layer.msg('请输入店铺编号', { icon: 5, time: 2000 });
                            return;
                        }

                        admin.ajax({
                            url: uri,
                            type: 'post',
                            data: {
                                url: $('#J_Title').val(),
                                store_id: $('#J_store_id').val()
                            },
                        }).done(function (res) {
                            if (res.code != '200') {
                                layer.msg(res.msg || '操作失败', { icon: 5, time: 2000 });
                                return;
                            }
                            layer.alert("保存成功", { icon: 6 }, function () {
                                layer.closeAll();
                            });
                        });
                    });
                }
            });
        });

        //链接关联
        $('.link-relation-form').on('click', function () {
            var uri = $(this).data('uri');

            layer.open({
                title: '表单信息',
                area: ['490px', '330px'],
                type: 1,
                content: '' +
                    '<form class="layui-form dialog-form" action="" id="J_Form">\n' +
                    '  <div class="dialog-form-item">' +
                    '                        <label class="dialog-form-label">关联类型</label>' +
                    '                        <div class="dialog-form-input">' +
                    '                            <input type="radio" lay-filter="rtype_even" name="r_type" value="0" checked title="指定小区" />' +
                    '                            <input type="radio" lay-filter="rtype_even" name="r_type" value="1" title="全部小区" />' +
                    '                        </div>' +
                    '                    </div>' +
                    '    <div class="dialog-form-item" id="form-ids">' +
                    '        <label class="dialog-form-label">小区ID</label>' +
                    '        <div class="dialog-form-input">' +
                    '            <input type="text" id="J_Title" name="ids" placeholder="请输入小区ID" maxlength="10" autocomplete="off" value="">' +
                    '            <p style="color:#999">多个小区ID使用英文逗号隔开</p>' +
                    '        </div>' +
                    '    </div>' +
                    '    <div class="dialog-form-item">' +
                    '        <label class="dialog-form-label">备注</label>' +
                    '        <div class="dialog-form-input">' +
                    '            <textarea name="remark" placeholder="请输入备注"></textarea>' +
                    '        </div>' +
                    '    </div>' +
                    '    <div class="dialog-form-item">' +
                    '        <label class="dialog-form-label">&nbsp;</label>' +
                    '        <div class="dialog-form-input">' +
                    '            <a href="javascript:;" id="J_Btn" class="dialog-form-btn">确定</a>' +
                    '            <a href="javascript:layer.closeAll();" class="dialog-form-hollowbtn">取消</a>' +
                    '        </div>' +
                    '    </div>' +
                    '</form>',
                success: function () {
                    form.render('radio');
                    form.on('radio(rtype_even)', function (data) {
                        if (data.value == 1) {
                            $('#form-ids').hide();
                        } else {
                            $('#form-ids').show();
                        }
                    });
                    $('#J_Btn').on('click', function () {

                        var title = $.trim($('#J_Title').val());
                        var stype = $("input[name='r_type']:checked").val();
                        console.log(stype)
                        if (stype == 0 && !title) {
                            layer.msg('请输入小区ID', { icon: 5, time: 2000 });
                            return;
                        }

                        admin.ajax({
                            url: uri,
                            type: 'post',
                            data: {
                                r_type: stype,
                                ids: $('#J_Title').val(),
                                remark: $('textarea[name=remark]').val()
                            },
                        }).done(function (res) {
                            if (res.code != '200') {
                                layer.msg(res.msg || '操作失败', { icon: 5, time: 2000 });
                                return;
                            }
                            layer.alert("保存成功", { icon: 6 }, function () {
                                layer.closeAll();
                            });
                        });
                    });
                }
            });
        });

        //标签添加
        $('.forum-tag-form').on('click', function () {
            var uri = $(this).data('uri');
            var title = $(this).data('title') || '';

            layer.open({
                title: '表单信息',
                area: ['520px', '240px'],
                type: 1,
                content: '' +
                    '<form class="layui-form dialog-form" action="" id="J_Form">\n' +
                    '    <div class="dialog-form-ite" id="form-store_id" style="width: 100%;height:40px;">\n' +
                    '        <label class="dialog-form-label">标签名称</label>\n' +
                    '        <div class="dialog-form-input">\n' +
                    '            <input type="text" id="J_Title" name="title" placeholder="请输入标签名称" style="width: 250px;" autocomplete="off" value="' + title + '">\n' +
                    '        </div>\n' +
                    '    </div>\n' +
                    '    <div class="dialog-form-item">\n' +
                    '        <label class="dialog-form-label">&nbsp;</label>\n' +
                    '        <div class="dialog-form-input">\n' +
                    '            <a href="javascript:;" id="J_Btn" class="dialog-form-btn">确定</a>\n' +
                    '            <a href="javascript:layer.closeAll();" class="dialog-form-hollowbtn">取消</a>\n' +
                    '        </div>\n' +
                    '    </div>\n' +
                    '</form>',
                success: function () {
                    $('#J_Btn').on('click', function () {

                        var title = $.trim($('#J_Title').val());
                        if (!title) {
                            layer.msg('请输入标签名称', { icon: 5, time: 2000 });
                            return;
                        }

                        admin.ajax({
                            url: uri,
                            type: 'post',
                            data: {
                                title: $('#J_Title').val()
                            }
                        }).done(function (res) {
                            if (res.code != '200') {
                                layer.msg(res.msg || '操作失败', { icon: 5, time: 2000 });
                                return;
                            }
                            layer.alert("保存成功", { icon: 6 }, function () {
                                layer.closeAll();
                                window.location.reload()
                            });
                        });
                    });
                }
            });
        });

        //删除宣传图
        $('.activity-adv-section').on('click', '.activity-btn-delete', function () {
            $(this).parents('.layui-form-section-item').remove();
        });

        //宣传图
        $('#btn_add_adv').click(function () {
            var index = $('.layui-form-section-item').length;
            var id_name = "J_TitleImgAdv_" + index;
            var _html = ['<div class="layui-form-section-item">',
                '                            <div class="layui-form-item">',
                '                                <label class="layui-form-label">图片标题：</label>',
                '                                <div class="layui-input-inline layui-input-inline-long">',
                '                                    <input name="ad_title[]" type="text" maxlength="30" placeholder="请输入图片标题" ',
                '                                           value="" class="layui-input">',
                '                                </div>',
                '                                <button type="button" class="layui-btn layui-btn-sm activity-btn-delete"><i class="layui-icon"></i></button>',
                '                            </div>',
                '                            <div class="layui-form-item">',
                '                                <label class="layui-form-label">图片：</label>',
                '                                <div class="layui-input-inline">',
                '                                    <button type="button" class="layui-btn J_UploadImgBtn" id="' + id_name + '" data-itype="activity_ad">',
                '                                        <i class="layui-icon">&#xe67c;</i>更改图片',
                '                                    </button>',
                '                                    <img src="/static/images/default.png" data-img="/static/images/default.png" class="J_ViewImg upload-img">',
                '                                    <input name="adv_img[]" type="hidden" value="">',
                '                                </div>',
                '                            </div>',
                '                        </div>'].join("");

            $('.activity-adv-section').append(_html);

            //上传组件注册
            uploadRegisterFn($('#' + id_name));
        });

        //是否开启返利
        form.on('switch(reward)', function (data) {
            if (!data.elem.checked) {
                $('.reward-section').addClass('hide');
                $(this).val(0);
            } else {
                $('.reward-section').removeClass('hide');
                $(this).val(1);
            }

        });


        //生成的图片上传组件
        var uploadRegisterFn = function (obj) {
            $this = $(obj);
            upload.render({
                elem: $this,
                url: '/admin/uploadify/upload?img_type=' + $this.attr('data-itype'),
                size: 1024,
                field: 'Filedata',
                acceptMime: 'image/*',
                before: function () {
                    layer.load();
                },
                done: function (res) {
                    layer.closeAll('loading');

                    if (res.code == '200') {
                        $this.siblings('img').removeClass('hide').attr({ 'src': res.msg }).data('img', res.msg);
                        $this.siblings('input[type="hidden"]').val(res.msg);
                    } else {
                        layer.msg(res.msg || '操作失败', { icon: 5, time: 2000 });
                    }
                },
                error: function () {
                    layer.closeAll('loading');
                    layer.msg('网络错误', { icon: 5, time: 2000 });
                }
            });
        };

        //窗口弹窗方法
        var dialogFn = function (data) {
            $this = $('#J_Pzbox');
            upload.render({
                elem: '#J_Upload',
                url: '/admin/uploadify/upload?img_type=' + $this.data('itype'),
                size: 1024,
                field: 'Filedata',
                acceptMime: 'image/*',
                before: function () {
                    layer.load();
                },
                done: function (res) {
                    layer.closeAll('loading');

                    if (res.code == '200') {
                        _img = '<span class="J_ImgItem">' +
                            '   <input type="hidden" name="img" value="' + res.msg + '"/>' +
                            '   <a href="javascript:;" class="J_ViewImg" data-img="' + res.msg + '">图片</a>' +
                            '   <a href="javascript:;" data-img="' + res.msg + '" class="layui-icon layui-icon-close del J_Del"></a>' +
                            '</span>';
                        $('#J_Pzbox').append(_img);
                        $('#J_Upload').prop('disabled', true);
                    } else {
                        layer.msg(res.msg || '操作失败', { icon: 5, time: 2000 });
                    }
                },
                error: function () {
                    layer.closeAll('loading');
                    layer.msg('网络错误', { icon: 5, time: 2000 });
                }
            });
        };

    });

});