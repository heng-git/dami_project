$(function(){
	baseApp.init(); //需要加;号
	$(window).resize(function(){//当窗口重置的时候重新计算高度
		baseApp.resizeIframe();
	})
})
var baseApp={
	init:function(){
		this.initAside()
		this.confirmDelete()
		this.resizeIframe()
		this.changeStatus()
	},
	initAside:function(){
		$('.aside h4').click(function(){
			$(this).siblings('ul').slideToggle();
		})
	},
	//设置iframe的高度
	resizeIframe:function(){
		$("#rightMain").height($(window).height()-80)//标签选择
	},
	// 删除提示
	confirmDelete:function(){
		$(".delete").click(function(){ //类别选择
			var flag=confirm("您确定要删除吗?")
			return flag
		})
	},
	changeStatus:function (){
		$(".chStatus").click(function(){
			var id=$(this).attr("data-id")//data-id传过来的数据
			var table=$(this).attr("data-table")
			var field=$(this).attr("data-field")
			var el =$(this)

			$.get("/admin/changeStatus",{id:id,table:table,field:field},function (response){ //第二个参数为连通url请求发送给服务器的数据
				if (response.success){  //相当于传来修改成功的消息的时候  就把当前点击的图片换了  与数据库中的修改其实并无关联
					// 刷新时数据库每次还是根据status按照原来未修改的src进行显示
						if (el.attr("src").indexOf("yes")!=-1){
							el.attr("src","/static/admin/images/no.gif")
						}else{
							el.attr("src","/static/admin/images/yes.gif")
						}
				}
			})
		})
		$(".chSpanNum").click(function (){
			// 1、获取el 以及el里面的属性值  注意$(this)指代当前对象span
			var id = $(this).attr("data-id")
			var table = $(this).attr("data-table")
			var field = $(this).attr("data-field")
			var num = $(this).html().trim()//获得span里面的内容并且去掉空格
			var spanEl = $(this)
			//2、创建一个input的dom节点   var input=$("<input value='' />");
			var input = $("<input style='width:60px'  value='' />");
			// 3、把input放在el里面  <span><input/></span>  替代原本span里面的值
			$(this).html(input);
			//4、让input获取焦点并给该焦点赋初值    $(input).trigger('focus').val(val);
			$(input).trigger("focus").val(num);
			// 5、点击input的时候阻止冒泡  即当前的input框并没有把前面的span框给覆盖掉  点击input框的时候也会触发span的点击事件
			$(input).click(function (e) {
				e.stopPropagation();
			})
			//6、鼠标离开的时候(blur)给span赋值,并触发ajax请求
			$(input).blur(function () {
				var inputNum = $(this).val()//这里的this指的是调用这个函数的input
				spanEl.html(inputNum)
				// 触发ajax请求  即先修改在html页面的显示 再修改数据库中的数据
				$.get("/admin/changeNum", { id: id, table: table, field: field, num: inputNum }, function (response) {//同时修改数据库中的信息
					console.log(response)
				}
				)
			})
		})
	}

}