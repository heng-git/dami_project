package main

import (
	"github.com/gin-contrib/cors"
	"html/template"
	"xiaomi_project/models"
	"xiaomi_project/routers"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	// 创建一个默认的路由引擎
	r := gin.Default()
	//配置gin允许跨域请求
	r.Use(cors.Default())
	//自定义模板函数  注意要把这个函数放在加载模板前
	r.SetFuncMap(template.FuncMap{ //这里只调用不执行
		"UnixToTime": models.UnixToTime,
		"Str2Html":   models.Str2Html,
		"FormatImg":  models.FormatImg,
		"Sub":        models.Sub,
		"Substr":     models.Substr,
		"FormatAttr": models.FormatAttr,
		"Mul":        models.Mul,
	})

	//加载并渲染templates文件夹下所有文件夹下的所有文件模板 放在配置路由前面
	r.LoadHTMLGlob("templates/**/**/*")
	//配置静态web目录   第一个参数表示路由, 第二个参数表示映射的目录
	r.Static("/static", "./static")

	// 创建基于 cookie 的存储引擎，secret11111 参数是用于加密的密钥
	store := cookie.NewStore([]byte("secret111"))
	//配置session的中间件 store是前面创建的存储引擎，我们可以替换成其他存储引擎
	r.Use(sessions.Sessions("mysession", store))

	routers.AdminRoutersInit(r)

	routers.ApiRoutersInit(r)

	routers.DefaultRoutersInit(r)

	r.Run()
}
