(function($){

    var app={
        init:function(){
    
            this.initSwiper();

            this.initNavSlide();
            this.initProductContentTab();
            this.initProductContentColor();
        },
        initSwiper:function(){    
            new Swiper('.swiper-container', {
                loop : true,
                navigation: {
                  nextEl: '.swiper-button-next',
                  prevEl: '.swiper-button-prev'                 
                },
                pagination: {
                    el: '.swiper-pagination',
                    clickable :true
                }
                
            });
        },
        initNavSlide:function(){
             $("#nav_list>li").hover(function(){

                $(this).find('.children-list').show();
             },function(){
                $(this).find('.children-list').hide(); 
             })          

        },
        initProductContentTab:function(){
            $(function () {
                $('.detail_info .detail_info_item:first').addClass('active');
                $('.detail_list li:first').addClass('active');
                $('.detail_list li').click(function () {
                    var index = $(this).index();
                    $(this).addClass('active').siblings().removeClass('active');
                    $('.detail_info .detail_info_item').removeClass('active').eq(index).addClass('active');

                })
            })
        },
        initProductContentColor:function(){
            var _that=this;//此this代表的是上面的app对象
            $("#color_list .banben").first().addClass("active");//打开详情页时默认显示的是第一个颜色的数据
            $("#color_name").html($("#color_list .active .yanse").html())//获取color_list下面颜色类中的值并渲染到color_name中
            $("#color_list .banben").click(function(){
                $(this).addClass("active").siblings().removeClass("active");//此处的this代表的是该dom行为
                $("#color_name").html($("#color_list .active .yanse").html())
                var goods_id=$(this).attr("goods_id")
                var color_id=$(this).attr("color_id")

                $.get("/product/getImgList",{"goods_id":goods_id,"color_id":color_id},function(response){
                    console.log(response)
                    if(response.success==true){
                        var swiperStr=""
                        for (var i = 0; i < response.result.length; i++) {
                            swiperStr += '<div class="swiper-slide"><img src="' + response.result[i].img_url + '"> </div>';
                            // console.log(swiperStr)
                        }
                        $("#item_focus").html(swiperStr)
                        _that.initSwiper()//代表使用app对象的其中一个方法
                    }
                })
            })
        },
    }   
    
    $(function(){
    
    
        app.init();
    })

    

})($)
