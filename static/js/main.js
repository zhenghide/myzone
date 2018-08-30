//登陆
function login() {
    var userName = $("#userNameL").val();
    var password = $("#passwordL").val();
    if(userName == ""){
        $(".error").html("用户名不能为空！")
    }else if(userName != "" && password == ""){
        $(".error").html("密码不能为空！")
    }else if(userName != "" && password != ""){
        $(".error").html("")
        $.ajax({
            url:'/login',
            type : "POST",
            cache:false,//false是不缓存，true为缓存
            async:true,//true为异步，false为同步
            data:{
                userName:userName,
                password:password
            },
            beforeSend:function(){
                //请求前
            },
            success:function(result){
                //请求成功时
                if(result.error != null){
                    $(".errorL").html(result.error);
                }else if(result.error == null && result.msg != null){
                    window.location.href = "/home";
                }
            },
            complete:function(){
                //请求结束时
            },
            error:function(){
                //请求失败时
            }
        })
    }
}

//注册
function register() {
    var nickName = $("#nickName").val();
    var userName = $("#userNameR").val();
    var password = $("#passwordR").val();
    var phone = $("#phone").val();
    var email = $("#email").val();
    if(userName == ""){
        $(".error").html("用户名不能为空！")
    }else if(userName != "" && password == ""){
        $(".error").html("密码不能为空！")
    }else if(userName != "" && password != ""){
        $(".error").html("")
        $.ajax({
            url:'/register',
            type : "POST",
            cache:false,//false是不缓存，true为缓存
            async:true,//true为异步，false为同步
            data:{
                nickName:nickName,
                userName:userName,
                password:password,
                phone:phone,
                email:email
            },
            beforeSend:function(){
                //请求前
            },
            success:function(result){
                //请求成功时
                if(result.error != null){
                    $(".errorR").html(result.error);
                }else if(result.error == null && result.msg != null){
                    window.location.href = "/home";
                }
            },
            complete:function(){
                //请求结束时
            },
            error:function(){
                //请求失败时
            }
        })
    }
}

//发表说说
function createSay() {
    var sayContent = $("#sayContent").val();
    if(sayContent != ""){
        $.ajax({
            url:'/say/create',
            type : "POST",
            cache:false,//false是不缓存，true为缓存
            async:true,//true为异步，false为同步
            data:{
                sayContent:sayContent
            },
            beforeSend:function(){
                //请求前
            },
            success:function(result){
                //请求成功时
                window.location.href = "/say";
            },
            complete:function(){
                //请求结束时
            },
            error:function(){
                //请求失败时
            }
        })
    }else{
        alert("要发表的说说不能为空！")
    }
}

//删除说说
function deleteSay(sayId) {
    $.ajax({
        url:'/say/delete',
        type : "POST",
        cache:false,//false是不缓存，true为缓存
        async:true,//true为异步，false为同步
        data:{
            sayId:sayId
        },
        beforeSend:function(){
            //请求前
        },
        success:function(result){
            //请求成功时
            if(result.msg != null){
                window.location.href = "/say";
            }
        },
        complete:function(){
            //请求结束时
        },
        error:function(){
            //请求失败时
        }
    })
}

//写日志
function createDiary() {
    var diaryTitle = $("#diaryTitle").val();
    var diaryContent = $('#diaryContent').summernote('code');

    alert(diaryContent);
    if(diaryContent != ""){
        $.ajax({
            url:'/diary/create',
            type : "POST",
            cache:false,//false是不缓存，true为缓存
            async:true,//true为异步，false为同步
            data:{
                diaryTitle:diaryTitle,
                diaryContent:diaryContent,
            },
            beforeSend:function(){
                //请求前
            },
            success:function(result){
                //请求成功时
                window.location.href = "/diary";
            },
            complete:function(){
                //请求结束时
            },
            error:function(){
                //请求失败时
            }
        })
    }else{
        alert("要发表的说说不能为空！")
    }
}

//删除说说
function deleteDiary(diaryId) {
    $.ajax({
        url:'/diary/delete',
        type : "POST",
        cache:false,//false是不缓存，true为缓存
        async:true,//true为异步，false为同步
        data:{
            diaryId:diaryId
        },
        beforeSend:function(){
            //请求前
        },
        success:function(result){
            //请求成功时
            if(result.msg != null){
                window.location.href = "/diary";
            }
        },
        complete:function(){
            //请求结束时
        },
        error:function(){
            //请求失败时
        }
    })
}

//创建相册
function CreatePhoto() {
    var photoName = $("#photoName").val();
    if(photoName != ""){
        $("#photoNameError").html("")
        $.ajax({
            url:'/photo/create',
            type : "POST",
            cache:false,//false是不缓存，true为缓存
            async:true,//true为异步，false为同步
            data:{
                photoName:photoName
            },
            beforeSend:function(){
                //请求前
            },
            success:function(result){
                //请求成功时
                window.location.href = "/photo";
            },
            complete:function(){
                //请求结束时
            },
            error:function(){
                //请求失败时
            }
        })
    }else{
        $("#photoNameError").html("相册名称不能为空！")
    }
}
