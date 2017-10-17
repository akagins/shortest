$(function(){
    $("#btn-signin").click(function() {
        $.ajax({
            url: $("#form-signin").attr("action"),
            type: $("#form-signin").attr("method"),
            dataType: 'json',
            // フォーム要素の内容をハッシュ形式に変換
            data: $("#form-signin").serializeArray(),
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