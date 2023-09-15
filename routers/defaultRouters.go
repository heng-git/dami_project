package routers

import (
	"github.com/gin-gonic/gin"
	"xiaomi_project/controllers/shopping"
	"xiaomi_project/middlewares"
)

func DefaultRoutersInit(r *gin.Engine) {
	defaultRouters := r.Group("/")
	{
		defaultRouters.GET("/", shopping.DefaultController{}.Index)
		defaultRouters.GET("/category:id", shopping.ProductController{}.Category)
		defaultRouters.GET("/detail", shopping.ProductController{}.Detail)
		defaultRouters.GET("/product/getImgList", shopping.ProductController{}.GetImgList)

		defaultRouters.GET("/cart", shopping.CartController{}.Get)
		defaultRouters.GET("/cart/addCart", shopping.CartController{}.AddCart)

		defaultRouters.GET("/cart/successTip", shopping.CartController{}.AddCartSuccess)

		defaultRouters.GET("/cart/decCart", shopping.CartController{}.DecCart)
		defaultRouters.GET("/cart/incCart", shopping.CartController{}.IncCart)

		defaultRouters.GET("/cart/changeOneCart", shopping.CartController{}.ChangeOneCart)
		defaultRouters.GET("/cart/changeAllCart", shopping.CartController{}.ChangeAllCart)
		defaultRouters.GET("/cart/delCart", shopping.CartController{}.DelCart)

		defaultRouters.GET("/pass/login", shopping.PassController{}.Login)
		defaultRouters.GET("/pass/captcha", shopping.PassController{}.Captcha)

		defaultRouters.GET("/pass/registerStep1", shopping.PassController{}.RegisterStep1)
		defaultRouters.GET("/pass/registerStep2", shopping.PassController{}.RegisterStep2)
		defaultRouters.GET("/pass/registerStep3", shopping.PassController{}.RegisterStep3)
		defaultRouters.GET("/pass/sendCode", shopping.PassController{}.SendCode)
		defaultRouters.GET("/pass/validateSmsCode", shopping.PassController{}.ValidateSmsCode)
		defaultRouters.POST("/pass/doRegister", shopping.PassController{}.DoRegister)
		defaultRouters.POST("/pass/doLogin", shopping.PassController{}.DoLogin)
		defaultRouters.GET("/pass/loginOut", shopping.PassController{}.LoginOut)

		defaultRouters.GET("/alipay", middlewares.InitUserAuthMiddleware, shopping.AlipayController{}.Alipay)
		defaultRouters.POST("/alipayNotify", shopping.AlipayController{}.AlipayNotify)
		defaultRouters.GET("/alipayReturn", middlewares.InitUserAuthMiddleware, shopping.AlipayController{}.AlipayReturn)

		//判断用户权限
		defaultRouters.GET("/buy/checkout", middlewares.InitUserAuthMiddleware, shopping.BuyController{}.Checkout) //判断用户权限
		defaultRouters.POST("/buy/doCheckout", middlewares.InitUserAuthMiddleware, shopping.BuyController{}.DoCheckout)
		defaultRouters.GET("/buy/pay", middlewares.InitUserAuthMiddleware, shopping.BuyController{}.Pay)
		defaultRouters.GET("/buy/orderPayStatus", middlewares.InitUserAuthMiddleware, shopping.BuyController{}.OrderPayStatus)

		defaultRouters.POST("/address/addAddress", middlewares.InitUserAuthMiddleware, shopping.AddressController{}.AddAddress)
		defaultRouters.POST("/address/editAddress", middlewares.InitUserAuthMiddleware, shopping.AddressController{}.EditAddress)
		defaultRouters.GET("/address/changeDefaultAddress", middlewares.InitUserAuthMiddleware, shopping.AddressController{}.ChangeDefaultAddress)
		defaultRouters.GET("/address/getOneAddressList", middlewares.InitUserAuthMiddleware, shopping.AddressController{}.GetOneAddressList)

		defaultRouters.GET("/wxpay", middlewares.InitUserAuthMiddleware, shopping.WxpayController{}.Wxpay)
		defaultRouters.POST("/wxpay/notify", shopping.WxpayController{}.WxpayNotify)

		defaultRouters.GET("/user", middlewares.InitUserAuthMiddleware, shopping.UserController{}.Index)
		defaultRouters.GET("/user/order", middlewares.InitUserAuthMiddleware, shopping.UserController{}.OrderList)
		defaultRouters.GET("/user/orderinfo", middlewares.InitUserAuthMiddleware, shopping.UserController{}.OrderInfo)

		defaultRouters.GET("/search", shopping.SearchController{}.Index)
		defaultRouters.GET("/search/getOne", shopping.SearchController{}.GetOne)
		defaultRouters.GET("/search/addGoods", shopping.SearchController{}.AddGoods)
		defaultRouters.GET("/search/updateGoods", shopping.SearchController{}.UpdateGoods)
		defaultRouters.GET("/search/deleteGoods", shopping.SearchController{}.DeleteGoods)
		defaultRouters.GET("/search/query", shopping.SearchController{}.Query)
		defaultRouters.GET("/search/filterQuery", shopping.SearchController{}.FilterQuery)
		defaultRouters.GET("/search/goodsList", shopping.SearchController{}.PagingQuery)
		defaultRouters.GET("/search/pagingQuery", shopping.SearchController{}.PagingQuery)
	}
}
