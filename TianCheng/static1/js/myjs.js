function sign() {
    $.ajax({
        type : 'GET',
        url : 'http://127.0.0.1:8082/sign1', //服务路径
        async : false,
        success : function (response) {
            if(response.code == 200){
                window.location.href ="http://127.0.0.1:8082/sign2";
            }else{
                alert(response.code);
            }
        },
        error : function (){
            alert('服务器异常！');
        }
    });
}
function login() {
    $.ajax({
        type : 'GET',
        url : 'http://127.0.0.1:8082/login1', //服务路径
        async : false,
        success : function (response) {
            if(response.code == 200){
                window.location.href ="http://127.0.0.1:8082/login2";
            }else{
                alert(response.code);
            }
        },
        error : function (){
            alert('服务器异常！');
        }
    });
}
