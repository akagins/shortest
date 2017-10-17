$(function(){
    $("#btn-signup").click(function() {
        $.ajax({
            url: $("#form-signup").attr("action"),
            type: $("#form-signup").attr("method"),
            dataType: 'json',
            // フォーム要素の内容をハッシュ形式に変換
            data: $("#form-signup").serializeArray(),
            success: function(data){
                if(data.url != "") {
                    location.href = data.url
                } else {
                    $(".alert.alert-danger").text(data.message)
                    $(".alert.alert-danger").show()
                }
                console.log("success")
            },
            error: function(){
                console.log("error")
            }
        });
    });
});