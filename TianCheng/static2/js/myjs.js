function zhuce() {
    var itemForm = $('#itemForm');
    var formData = new FormData($("#itemForm")[0]);
    $.ajax({
        type : 'POST',
        url : 'http://127.0.0.1:8082/zhuce', //服务路径
        data:formData,
        dataType:"json",
        processData: false,
        contentType: false,
        async : false,
        success : function (response) {
            if (response.code ==203) {
                if (response.username !="ok"){
                    alert(response.username)
                }
                if (response.email !="ok"){
                    alert(response.email)
                }
                if (response.phone !="ok"){
                    alert(response.phone)
                }
                if (response.password !="ok"){
                    alert(response.password)
                }
            }
            if (response.code ==200){
                window.location.href ="http://127.0.0.1:8082/login2";
            }

        },
        error : function (){
            alert('服务器异常！');
        }
    });
}
function login2() {
    var login = $('#login');
    var formData = new FormData($("#login")[0]);
    var text2=$("#text2").html();
    if (text2=="验证通过"){
        $.ajax({
            type: 'POST',
            url: 'http://127.0.0.1:8082/denglu', //服务路径
            data: formData,
            dataType: "json",
            processData: false,
            contentType: false,
            async: false,
            success: function (response) {
                if (response.code == 200) {
                    window.location.href ="http://127.0.0.1:8082/index?Authorization="+response.data.token;
                }
                else {
                    alert("用户名密码错误");
                }
            },
            error: function () {
                alert('服务器异常2！');
            }
        });
    }
    else {
        alert("请拖动滑块解锁")
    }
}
function ksxq(token,id) {
    $.ajax({
        type: 'POST',
        url: 'http://127.0.0.1:8082/ksxq?Authorization='+token+'&id='+id, //服务路径
        dataType: "json",
        processData: false,
        contentType: false,
        async: false,
        success: function (response) {
            if (response.code == 200) {
                $("#ksmc").html("考试名称："+response.kcname);
                $("#kssj").html("考试时间："+response.stime+"---"+response.etime);
                $("#kdcs").html("可答次数："+response.kdnumber);
                $("#dxt").html(response.dati+"大题，"+response.xiaoti+"小题");
                $("#zftg").html(response.zongf+"总分，"+response.pass+"通过");
                $("#fbr").html("发布人："+response.fabuid);
            }
            else {
                alert("用户名密码错误");
            }
        },
        error: function () {
            alert('服务器异常！');
        }
    });
}
function zj(token) {
    var id =$("#bm option:selected").val();
    $.ajax({
        type: 'POST',
        url: 'http://127.0.0.1:8082/zj?Authorization='+token+'&id='+id, //服务路径
        dataType: "json",
        processData: false,
        contentType: false,
        async: false,
        success: function (response) {
            if (response.code == 200) {
                $("#phones").val($("#phones").val()+response.phones)
            }
            else {
                alert("用户名密码错误");
            }
        },
        error: function () {
            alert('服务器异常！');
        }
    });
}
function jiaoj(token,id) {
    var login = $('#jiaojuan');
    var formData = new FormData($("#jioajuan")[0]);
    var t=$("#tm2").text();
        $.ajax({
            type: 'POST',
            url: 'http://127.0.0.1:8082/jiaojuan?Authorization='+token+'&examid='+id+"&tm="+t, //服务路径
            data: formData,
            dataType: "json",
            processData: false,
            contentType: false,
            async: false,
            success: function (response) {
                if (response.code == 200) {
                    window.location.href ="http://127.0.0.1:8082/index?Authorization="+response.token;
                }
                else {
                    alert("用户名密码错误");
                }
            },
            error: function () {
                alert('服务器异常！');
            }
        });
}
function sckc(id,token) {
    $.ajax({
        type: 'POST',
        url: 'http://127.0.0.1:8082/sckc?Authorization='+token+'&kcid='+id, //服务路径
        dataType: "json",
        processData: false,
        contentType: false,
        async: false,
        success: function (response) {
            if (response.code == 200) {
                window.location.href ="http://127.0.0.1:8082/kcsz?Authorization="+response.token;
            }
            else {
                alert("用户名密码错误");
            }
        },
        error: function () {
            alert('服务器异常！');
        }
    });
}
function scsj(id,token) {
    alert(id);
    $.ajax({
        type: 'POST',
        url: 'http://127.0.0.1:8082/scsj?Authorization='+token+'&sjid='+id, //服务路径
        dataType: "json",
        processData: false,
        contentType: false,
        async: false,
        success: function (response) {
            if (response.code == 200) {
                window.location.href ="http://127.0.0.1:8082/glsj?Authorization="+response.token;
            }
            else {
                alert("用户名密码错误");
            }
        },
        error: function () {
            alert('服务器异常！');
        }
    });
}
function xzbm(id,token) {
    var name=$("#xzbmname").val();
    $.ajax({
        type: 'POST',
        url: 'http://127.0.0.1:8082/xzbm?Authorization='+token+'&qyid='+id+'&name='+name, //服务路径
        processData: false,
        contentType: false,
        async: false,
        success: function (response) {
            if (response.code == 200) {
                window.location.href ="http://127.0.0.1:8082/index?Authorization="+response.token;
            }
            else {
                alert(response.code);
            }
        },
        error: function () {
            alert('服务器异常！');
        }
    });
}
function scbm(id,token) {
    var name=$("#scbmname").val();
    $.ajax({
        type: 'DELETE',
        url: 'http://127.0.0.1:8082/scbm?Authorization='+token+'&qyid='+id+'&name='+name, //服务路径
        dataType: "json",
        processData: false,
        contentType: false,
        async: false,
        success: function (response) {
            if (response.code == 200) {
                window.location.href ="http://127.0.0.1:8082/index?Authorization="+response.token;
            }
            else {
                alert("删除失败");
            }
        },
        error: function () {
            alert('服务器异常！');
        }
    });
}
function cxyg(token) {
    var bmname =$("#bm option:selected").val();
    var name=$("#ygname").val();
    window.location.href ="http://127.0.0.1:8082/cxyg?Authorization="+token+'&bmname='+bmname+'&name='+name;
}
function xg(id) {
    $("#userid").val(id)
}
function xxxg(token) {
 var id=$("#userid").val();
 var bm =$("#bmxg option:selected").val();
 var qx =$("#qxxg option:selected").val();
    $.ajax({
        type: 'PUT',
        url: 'http://127.0.0.1:8082/xxxg?Authorization='+token+'&id='+id+'&bm='+bm+"&qx="+qx, //服务路径
        dataType: "json",
        processData: false,
        contentType: false,
        async: false,
        success: function (response) {
            if (response.code == 200) {
                window.location.href ="http://127.0.0.1:8082/index?Authorization="+response.token;
            }
            else {
                alert("删除失败");
            }
        },
        error: function () {
            alert('服务器异常！');
        }
    });
}
function tjkc(token) {
    var kcname = $("#kcname").val();
    if (kcname==""){
        alert("请输入考场名称")
    }else {
        var maxqh=0;
        var val1=$('input[name="fzb"]:checked').val();
        if (val1==1){
            maxqh = $("#maxqh").val()
        }
        var val2=$('input[name="fzzt"]:checked').val();
        var val3=$('input[name="timu"]:checked').val();
        var val4=$("#datinum").val();
        $.ajax({
            type: 'PUT',
            url: 'http://127.0.0.1:8082/tjkc?Authorization='+token+'&kcname='+kcname+'&val1='+val1+'&val2='+val2+'&val3='+val3+'&val4='+val4+'&maxqh='+maxqh, //服务路径
            dataType: "json",
            processData: false,
            contentType: false,
            async: false,
            success: function (response) {
                if (response.code == 200) {
                    window.location.href ="http://127.0.0.1:8082/index?Authorization="+response.token;
                }
                else {
                    alert("新建失败");
                }
            },
            error: function () {
                alert('服务器异常！');
            }
        });
    }
}
function xgmm(token) {
    var mm1=$("#ymm").val();
    var mm2=$("#xmm").val();
    var mm3=$("#zcsr").val();
    if (mm2!=mm3){
        alert("密码不一致")
    }else {
        $.ajax({
            type: 'GET',
            url: 'http://127.0.0.1:8082/checkp?Authorization='+token+'&password='+mm1, //服务路径
            dataType: "json",
            processData: false,
            contentType: false,
            async: false,
            success: function (response) {
                if (response.code == 200) {
                    //window.location.href ="http://127.0.0.1:8082/index?Authorization="+response.token;
                    $.ajax({
                        type: 'PUT',
                        url: 'http://127.0.0.1:8082/xgmm?Authorization='+token+'&password='+mm2, //服务路径
                        dataType: "json",
                        processData: false,
                        contentType: false,
                        async: false,
                        success: function (response) {
                            if (response.code == 200) {
                                window.location.href ="http://127.0.0.1:8082/login2";
                            }
                            else {
                                alert("修改失败");
                            }
                        },
                        error: function () {
                            alert('服务器异常！');
                        }
                    });
                }
                else {
                    alert("密码不正确");
                }
            },
            error: function () {
                alert('服务器异常！');
            }
        });
    }
}
function ksby(id,token) {
    $("#t1").val(id);
    $("#t2").val(token);
}
function ksks() {
    var id=$("#t1").val();
    var token=$("#t2").val();
    var token2=$("#t3").val();
    window.location.href ="http://127.0.0.1:8082/ksks?id="+id+"&Authorization="+token+"&token2="+token2
}
function ws() {
    window.location.href ="http://127.0.0.1:8083/"
}
function zjcs(token) {
    $.ajax({
        type: 'GET',
        url: 'http://127.0.0.1:8082/zjcs?Authorization='+token, //服务路径
        dataType: "json",
        processData: false,
        contentType: false,
        async: false,
        success: function (response) {
            if (response.code == 200) {
                //window.open(response.url, '_blank').location;
            }
            else {
                alert("失败");
            }
        },
        error: function () {
            alert('服务器异常！');
        }
    });
}
function vip(token) {
    $.ajax({
        type: 'GET',
        url: 'http://127.0.0.1:8082/vip?Authorization='+token, //服务路径
        dataType: "json",
        processData: false,
        contentType: false,
        async: false,
        success: function (response) {
            if (response.code == 200) {
                //window.open(response.url, '_blank').location;
            }
            else {
                alert("失败");
            }
        },
        error: function () {
            alert('服务器异常！');
        }
    });
}
function fsyzm() {
    var email = $("#email").val();
    var e = checkemail(email);
   if (e=="success"){
       $.ajax({
           type: 'GET',
           url: 'http://127.0.0.1:8082/fsyzm?email='+email, //服务路径
           dataType: "json",
           processData: false,
           contentType: false,
           async: false,
           success: function (response) {
               if (response.code == 200) {
                  alert("已发送")
               }
               else if (response.code == 202) {
                   alert("5分钟之后再试");
               }
               else {
                   alert("发送失败")
               }
           },
           error: function () {
               alert('服务器异常！');
           }
       });
   }else {
       alert("邮箱不存在")
   }

}
function checkemail(email) {
    var s= "success";
    var f= "fail";
    var st ="";
    $.ajax({
        type: 'GET',
        url: 'http://127.0.0.1:8082/checkemail?email='+email, //服务路径
        dataType: "json",
        processData: false,
        contentType: false,
        async: false,
        success: function (response) {
            if (response.code == 200) {
               st="s"
            }
            else {
                st="f"
            }
        },
        error: function () {
            st="f"
        }
    });
    if (st=="s"){
        return s
    }else {
        return f
    }
}
function czmm() {
    var email = $("#email").val();
    var yzm = $("#yzm").val();
    if (email == "" || yzm == ""){
        alert("请输入信息")
    }else {
        $.ajax({
            type: 'GET',
            url: 'http://127.0.0.1:8082/czmm?email='+email+"&yzm="+yzm, //服务路径
            dataType: "json",
            processData: false,
            contentType: false,
            async: false,
            success: function (response) {
                if (response.code == 200) {
                    window.location.href ="http://127.0.0.1:8082/czm?email="+email
                }
                else {
                    alert("信息错误")
                }
            },
            error: function () {
                alert("服务器错误")
            }
        });
    }

}
function cmm() {
    var email = $("#email").val();
    var m1 = $("#p1").val();
    var m2 = $("#p2").val();
    if (m1!=m2){
        alert("密码不一致")
    }else{
        $.ajax({
            type: 'PUT',
            url: 'http://127.0.0.1:8082/cm?email='+email+"&password="+m1, //服务路径
            dataType: "json",
            processData: false,
            contentType: false,
            async: false,
            success: function (response) {
                if (response.code == 200) {
                    alert("重置成功");
                    window.location.href ="http://127.0.0.1:8082/login2";
                }
                else {
                    alert("重置失败")
                }
            },
            error: function () {
                alert("服务器错误")
            }
        });
    }
}
function glks(token,id) {
    $.ajax({
        type: 'GET',
        url: 'http://127.0.0.1:8082/glks?Authorization='+token+'&id='+id, //服务路径
        dataType: "json",
        processData: false,
        contentType: false,
        async: false,
        success: function (response) {
            if (response.code == 200) {
                window.location.href="http://127.0.0.1:8082/glks2?Authorization="+token+"&id="+id;
            }
            else {
                alert("考试已经开始，无法修改")
            }
        },
        error: function () {
            alert("服务器错误")
        }
    });
}
function scks(ksid,token,testid) {
    var person1 = new Object();
    person1.ksid=ksid;
    person1.testid=testid;
    var json_str = JSON.stringify(person1);
    $.ajax({
        type: 'DELETE',
        url: 'http://127.0.0.1:8082/scks?Authorization='+token, //服务路径
        data:json_str,
        dataType: "json",
        processData: false,
        contentType: false,
        async: false,
        success: function (response) {
            if (response.code == 200) {
                window.location.href="http://127.0.0.1:8082/ksgl?Authorization="+token;
            }
            else {
                alert("删除失败");
            }
        },
        error: function () {
            alert("服务器错误")
        }
    });
}
function tjks(token,testid) {
   var phones= $("#phones").val();
    var person1 = new Object();
    person1.phones=phones;
    person1.testid=testid;
    var json_str = JSON.stringify(person1);
    $.ajax({
        type: 'POST',
        url: 'http://127.0.0.1:8082/tjks?Authorization='+token, //服务路径
        data:json_str,
        dataType: "json",
        processData: false,
        contentType: false,
        async: false,
        success: function (response) {
            if (response.code == 200) {
                window.location.href="http://127.0.0.1:8082/ksgl?Authorization="+token;
            }
            else {
                alert("删除失败");
            }
        },
        error: function () {
            alert("服务器错误")
        }
    });
}
function pysj(token,testid) {
    $.ajax({
        type: 'GET',
        url: 'http://127.0.0.1:8082/pysj?Authorization='+token+'&id='+testid, //服务路径
        dataType: "json",
        processData: false,
        contentType: false,
        async: false,
        success: function (response) {
            if (response.code == 200) {
                window.location.href="http://127.0.0.1:8082/pysj2?Authorization="+token+"&id="+testid+"&li=0";
            }
            else {
                alert("考试未结束")
            }
        },
        error: function () {
            alert("服务器错误")
        }
    });
}
function jiaf(userid,id2,testid,token) {
    var g1 = $("#"+id2+"_1").val();//分数
    var g2 = $("#"+id2+"_2").val();//老师给的分数
        $.ajax({
            type: 'PUT',
            url: 'http://127.0.0.1:8082/jiaf?Authorization='+token+'&userid='+userid+'&testid='+testid+'&fs='+g2+'&g1='+g1, //服务路径
            dataType: "json",
            processData: false,
            contentType: false,
            async: false,
            success: function (response) {
                if (response.code == 200) {
                    $("#"+id2+"_3").hide()
                }
                else {
                    alert("分数不对")
                }
            },
            error: function () {
                alert("服务器错误")
            }
        });
}
function kssc(token,testid) {
    $.ajax({
        type: 'DELETE',
        url: 'http://127.0.0.1:8082/kssc?Authorization='+token+'&testid='+testid, //服务路径
        dataType: "json",
        processData: false,
        contentType: false,
        async: false,
        success: function (response) {
            if (response.code == 200) {
                window.location.href="http://127.0.0.1:8082/ksgl?Authorization="+response.token
            }
            else {
                alert("删除失败")
            }
        },
        error: function () {
            alert("服务器错误")
        }
    });
}
function drks(token) {
    var login = $('#ksxx');
    var formData = new FormData($("#ksxx")[0]);
    $.ajax({
        type: 'POST',
        url: 'http://127.0.0.1:8082/drks?Authorization='+token, //服务路径
        data: formData,
        dataType: "json",
        processData: false,
        contentType: false,
        async: false,
        success: function (response) {
            if (response.code == 200) {
                alert(response.phones);
                $("#phones").val(response.phones)
            }
            else {
                alert("失败");
            }
        },
        error: function () {
            alert('服务器异常！');
        }
    });
}