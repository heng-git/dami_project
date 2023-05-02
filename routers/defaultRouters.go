package routers

import (
	"github.com/gin-gonic/gin"
	"xiaomi_project/controllers/itying"
)

func DefaultRoutersInit(r *gin.Engine) {
	defaultRouters := r.Group("/")
	{
		defaultRouters.GET("/", itying.DefaultController{}.Index)
		defaultRouters.GET("/category:id", itying.ProductController{}.Category)
		defaultRouters.GET("/detail", itying.ProductController{}.Detail)
		defaultRouters.GET("/product/getImgList", itying.ProductController{}.GetImgList)
	}
}
