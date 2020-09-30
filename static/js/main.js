/* 返回顶部 */
$(document).ready(function() {
    $("body").append('<div class="goto_top"><a class="erweima" href="javascript:;"><span>微信二维码，扫一扫</span></a><a class="backindex" href="#" target="_blank"><span class="red_point"></span></a><a class="backtop" href="javascript:void(0);"></a></div>');
    $(".backtop").click(function() {
        $("html,body").animate({ scrollTop: 0 }, 1000);
    });
});

/* 左上轮播图 */
$(function() {
    if (typeof $(".banner").terseBanner === 'function') {
        $(".banner").terseBanner({
            useHover: true,
            auto: 3500,
            duration: 500,
            arrow: true
        });
    }
});

/* 左右竖栏跟随 */
jQuery(document).ready(function() {
    jQuery('.list_box_l,.list_right,.list_left_one').theiaStickySidebar({
        // Settings
        additionalMarginTop: 94
    });
});

/* 投稿开关 */
function closepopit() {
    $('.contribute-pop').hide();
    return false;
}

function showpopit() {
    $('.contribute-pop').show();
    return false;
}
document.writeln("<div class=\'contribute-pop\' style=\'display: none;\'>");
document.writeln("<div class=\'contribute-pop-content\'>");
document.writeln("<p>投稿邮箱：tougao@caijinglou.com</p>");
document.writeln("<p>一经录用会有专人和您联系</p>");
document.writeln("<p style=\'margin-top:10px\'>开通楼主号、入驻财经楼请咨询</p>");
document.writeln("<p>微信客服：炎炎燚燚</p>");
document.writeln("<p><img src=\'http://style.caijinglou.com/pc/images/cjl-wx-kefu.jpg\' alt=\'炎炎燚燚\'/></p>      ");
document.writeln("<button onclick=\'closepopit();\'>我知道了</button>");
document.writeln("</div>");
document.writeln("</div>");

/* 顶部导航悬浮 */
$(function() {
    $(window).scroll(function() {
        if ($(window).scrollTop() >= 40) {
            $("div.header").addClass("fixedNav");
        } else {
            $("div.header").removeClass("fixedNav");
        }
    });
});

/* 首页排行切换 */
$(document).ready(function() {
    $('#div_hottabs').children('span').bind("mouseenter", function() {
        changHot($(this).attr("idx"));
        $(this).parent().children().removeClass("cur");
        $(this).addClass("cur");
    });
});

function changHot(idx) {
    var ul = $('#div_hotlist').find('ul[idx=' + idx + ']');
    ul.siblings().hide();
    ul.show();
}

/* 内容页打印文章 */

(function($) {

    var printAreaCount = 0;

    $.fn.printArea = function() {

        var ele = $(this);

        var idPrefix = "printArea_";

        removePrintArea(idPrefix + printAreaCount);

        printAreaCount++;

        var iframeId = idPrefix + printAreaCount;

        var iframeStyle = 'position:absolute;width:0px;height:0px;left:-500px;top:-500px;';

        iframe = document.createElement('IFRAME');

        $(iframe).attr({

            style: iframeStyle,

            id: iframeId

        });


        document.body.appendChild(iframe);

        var doc = iframe.contentWindow.document;

        $(document).find("link").filter(function() {
            return $(this).attr("rel").toLowerCase() == "stylesheet";

        }).each(

            function() {

                doc.write('<link type="text/css" rel="stylesheet" href="' +
                    $(this).attr("href") + '" >');

            });

        doc.write('<div class="' + $(ele).attr("class") + '">' + $(ele).html()

            +
            '</div>');

        doc.close();

        var frameWindow = iframe.contentWindow;

        frameWindow.close();

        frameWindow.focus();

        frameWindow.print();

    }

    var removePrintArea =

        function(id) {

            $("iframe#" + id).remove();

        };
})(jQuery);

$(document).ready(function() {
    $(".print").click(function() {
        $(".content").printArea();
    });
});

/*获取当前日期时间*/

var myDate = new Date;
var year = myDate.getFullYear(); //获取当前年
var mon = myDate.getMonth() + 1; //获取当前月
var date = myDate.getDate(); //获取当前日
// var h = myDate.getHours();//获取当前小时数(0-23)
// var m = myDate.getMinutes();//获取当前分钟数(0-59)
// var s = myDate.getSeconds();//获取当前秒
var week = myDate.getDay();
var weeks = ["星期日", "星期一", "星期二", "星期三", "星期四", "星期五", "星期六"];
$("#time").html(year + "年" + mon + "月" + date + "日" + weeks[week]);



/*切换颜色*/
$(function () {
    var $li = $("#skin span");  //查找到元素
    $li.click(function () {   //给元素添加事件
        console.log(this.id);
        switchSkin(this.id);//调用函数

    });
    //保存Cookie完毕以后就可以通过Cookie来获取当前的皮肤了
    // var cookie_skin = $.cookie("MyCssSkin");     //获取Cookie的值
    // if (cookie_skin) {                          //如果确实存在Cookie
    //      switchSkin(cookie_skin);     //执行
    // }
});
function switchSkin(skinName) {   
    $("#" + skinName).addClass("active").siblings().removeClass("active");  //去掉其他同辈<li>元素的选中
    $("#cssfile").attr("href", "css/skin/" + skinName + ".css"); //设置不同皮肤
    //$.cookie("MyCssSkin", skinName, { path: '/', expires: 10 });  //保存Cookie
}