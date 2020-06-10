function login() {
    var username = $("#username").val();
    var password = $("#password").val();
    if (username !="" && password!=""){
        $.ajax({
            type: 'GET',
            url: 'http://127.0.0.1:8084/login?'+'&username='+username+'&password='+password, //服务路径
            dataType: "json",
            processData: false,
            contentType: false,
            async: false,
            success: function (response) {
                if (response.code == 200) {
                    window.location.href="http://127.0.0.1:8084/index?Authorization="+response.token;
                }
                else {
                    alert("失败");
                }
            },
            error: function () {
                alert("服务器错误")
            }
        });
    }
    else{
        alert("请输入用户名和密码");
    }
}
function czyh(token) {
    var name = $("#czname").val();
    if (name==""){
        alert("请输入要查找的用户信息")
    }else{
        var ysy = $("#ysy option:selected").val();
        window.location.href='http://127.0.0.1:8084/findusers?Authorization='+token+'&name='+name+'&ysy='+ysy
    }
}
function czqy(token) {
    var qyname = $("#qyname").val();
    window.location.href='http://127.0.0.1:8084/qygl?Authorization='+token+'&qyname='+qyname
}
function xzqy(token) {
    var qymc = $("#qymc").val();
    var glyname = $("#glyname").val();
    var glymm = $("#glymm").val();
    var glymmzc = $("#glymmzc").val();
    var glyphone = $("#glyphone").val();
    var glyemail = $("#glyemail").val();
    var glynumber = $("#glynumber").val();
    if (qymc=="" || glyname=="" || glymm=="" || glymmzc=="" || glyphone=="" || glyemail=="" || glynumber=="" || glymm!=glymmzc ){
        alert("请输入完整信息或密码不一致")
    }else{
        $.ajax({
            type: 'GET',
            url: 'http://127.0.0.1:8084/xzqycheck?Authorization='+token+'&qymc='+qymc+'&glyname='+glyname+'&glyphone='+glyphone+'&glymm='+glymm+'&glynumber='+glynumber, //服务路径
            dataType: "json",
            processData: false,
            contentType: false,
            async: false,
            success: function (response) {
                if (response.code == 200) {
                    window.location.href="http://127.0.0.1:8084/xzqy?Authorization="+token+"&qymc="+qymc+"&glyname="+glyname+"&glyphone="+glyphone+"&glyemail="+glyemail+"&glyemail="+glyemail+"&glynumber="+glynumber;
                }
                else {
                    alert("信息重复");
                }
            },
            error: function () {
                alert("服务器错误")
            }
        });
    }
}

function tznr(id,token) {
    $.ajax({
        type: 'GET',
        url: 'http://127.0.0.1:8084/tznr?Authorization='+token+'&id='+id, //服务路径
        dataType: "json",
        processData: false,
        contentType: false,
        async: false,
        success: function (response) {
            if (response.code == 200) {
                $("#title1").text(response.tz.Title);
                $("#text1").text(response.tz.Text);
            }
            else {
                alert("信息重复");
            }
        },
        error: function () {
            alert("服务器错误")
        }
    });
}
function bjtz(id,token) {
    $.ajax({
        type: 'GET',
        url: 'http://127.0.0.1:8084/tznr?Authorization='+token+'&id='+id, //服务路径
        dataType: "json",
        processData: false,
        contentType: false,
        async: false,
        success: function (response) {
            if (response.code == 200) {
                $("#title2").val(response.tz.Title);
                $("#text2").val(response.tz.Text);
                $("#tzid").val(response.tz.Id);
            }
            else {
                alert("信息重复");
            }
        },
        error: function () {
            alert("服务器错误")
        }
    });
}
function bjtz2(token) {
    var title = $("#title2").val();
    var text = $("#text2").val();
    var tzid = $("#tzid").val();
    window.location.href="http://127.0.0.1:8084/bjtz?Authorization="+token+"&title="+title+"&text="+text+"&tzid="+tzid
}
function xztz(token) {
    var title = $("#title3").val();
    var text = $("#text3").val();
    window.location.href="http://127.0.0.1:8084/xztz?Authorization="+token+"&title="+title+"&text="+text
}
function cztz(token) {
    var title = $("#gjz2").val();
    window.location.href="http://127.0.0.1:8084/tzgl?Authorization="+token+"&title="+title
}
function lh(token,id) {
    $.ajax({
        type: 'GET',
        url: 'http://127.0.0.1:8084/lh?Authorization='+token+'&id='+id, //服务路径
        dataType: "json",
        processData: false,
        contentType: false,
        async: false,
        success: function (response) {
            if (response.code==200){
                alert("已禁止")
            }
        },
        error: function () {
            alert("服务器错误")
        }
    });
}
function jc(token,id) {
    $.ajax({
        type: 'GET',
        url: 'http://127.0.0.1:8084/jc?Authorization='+token+'&id='+id, //服务路径
        dataType: "json",
        processData: false,
        contentType: false,
        async: false,
        success: function (response) {
            if (response.code==200){
                alert("已解除")
            }
        },
        error: function () {
            alert("服务器错误")
        }
    });
}
function sctz(id,token) {
    $.ajax({
        type: 'GET',
        url: 'http://127.0.0.1:8084/sctz?Authorization='+token+'&id='+id, //服务路径
        dataType: "json",
        processData: false,
        contentType: false,
        async: false,
        success: function (response) {
            if (response.code==200){
                window.location.href="http://127.0.0.1:8084/tzgl?Authorization="+token
            }
        },
        error: function () {
            alert("服务器错误")
        }
    });
}