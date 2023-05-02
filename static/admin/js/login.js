$(function (){ //$定义一个方法：在页面完全加载前而DOM（浏览器自带对象）加载后响应事件
    app.init()
})
var app={
    init:function (){
        this.getCaptcha()
        this.captchaImgChage()
    },
    getCaptcha:function (){
        //添加?t= +math.random() 是为了解决浏览器缓存的问题  每次请求添加一个随机的时间
        $.get("/admin/captcha?t="+Math.random(),function (response){ //访问captcha  返回的是一个gin.H(即map类型)的数据
            console.log(response)
            $("#captchaId").val(response.captchaId)//在这里$充当选择器   给id为#captchaId的标签value赋值为response.captchaId
            $("#captchaImg").attr("src",response.captchaImage)//给src属性赋值为captchaImage
        })  //请求当前域名下的/admin/captcha路径
    },
    captchaImgChage:function(){
        var that=this;
        $("#captchaImg").click(function(){
            that.getCaptcha()
        })
    }
}

